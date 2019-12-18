package sdproject

import (
	"fmt"
	"math"
	"sdproject/protos"
	"sync"

	"github.com/hashicorp/raft"
	"google.golang.org/grpc"
)

type NodeData struct {
	Id            int
	RaftID        int
	Address       string
	Parent        string
	RaftAddress   string
	RaftDirectory string
}

type NodeMtx struct {
	StorageMtx     sync.RWMutex
	FingerTableMtx sync.RWMutex
	PredecessorMtx sync.RWMutex
	SnapshotMtx    sync.RWMutex
	PoolMtx        sync.RWMutex
}

type Node struct {
	*protos.Node
	protos.ChordServer
	NodeMtx
	Parent        string
	Predecessor   *protos.Node
	Storage       map[int64]string
	FingerTable   []*protos.Node
	Pool          map[string]*GrpcConn
	GRPCServer    *grpc.Server
	Replicas      []*protos.Node
	RaftID        int
	Raft          *raft.Raft
	RaftAddress   string
	RaftDirectory string
	Leader        *protos.Node
	StopNode      chan struct{}
}

func NewNode(data *NodeData, isReplica bool) (*Node, error) {
	var err error
	NewConfig()
	node := &Node{}
	node.Node = new(protos.Node)
	node.Id = int64(data.Id)
	node.Parent = data.Parent
	node.Address = data.Address
	node.RaftID = data.RaftID
	node.RaftAddress = data.RaftAddress
	node.RaftDirectory = data.RaftDirectory
	node.Replicas = []*protos.Node{}
	node.StopNode = make(chan struct{})
	node.Pool = make(map[string]*GrpcConn)
	node.FingerTable = make([]*protos.Node, GetIntEnv("CHORD_SIZE"))
	node.Storage = make(map[int64]string)

	node.Raft, err = node.newRaftNode(isReplica)
	if err != nil {
		NewTracer("error", "NewNode::newRaftNode", err.Error())
		return nil, err
	}

	if isReplica {
		leaderNode := &protos.Node{Address: data.Parent}
		joinNode := &protos.Node{Id: int64(node.Id), Address: data.Address}
		joinRaftNode := &protos.Node{Id: int64(node.RaftID), Address: data.RaftAddress}
		node.Leader, err = node.JoinRaftGRPC(leaderNode, &protos.MultipleNodes{ChordNode: joinNode, RaftNode: joinRaftNode})
		if err != nil {
			NewTracer("error", "NewNode::JoinRaftGRPC", err.Error())
			return nil, err
		}
	} else {
		<-node.Raft.LeaderCh()
		node.Leader = node.Node
	}

	listen, err := StartTCPServer(data.Address)
	if err != nil {
		return nil, err
	}

	node.GRPCServer = grpc.NewServer()
	protos.RegisterChordServer(node.GRPCServer, node)

	go node.GRPCServer.Serve(listen)
	go node.asyncFixRaftLeader()

	if !isReplica {
		node.createChordOrJoinNode(data.Parent)
	}

	return node, nil
}

func (node *Node) StartAsyncProcess() {
	go node.asyncStabilize()
	go node.asyncFixFingerTable()
	go node.asyncFixStorage()
}

func (node *Node) fixStorage() {
	for key, value := range node.Storage {
		closestNode, _ := node.findSuccessor(key)

		if closestNode.Id == node.Id {
			return
		}

		node.StorageSetGRPC(closestNode, key, value)
		node.applyDelete(key)
	}
}

func (node *Node) LeaveNode() {
	NewTracer("info", "LeaveNode", "Leave node...")

	node.GRPCServer.Stop()

	if len(node.Replicas) > 0 {
		replicaSuccessor := node.Replicas[0]

		_, err := node.JoinGRPC(replicaSuccessor, &protos.Node{Address: node.Parent})
		if err != nil {
			NewTracer("error", "LeaveNode::JoinGRPC", err.Error())
			return
		}

		NewTracer("info", "LeaveNode::JoinGRPC", "Add replica to chord!")

		for _, conn := range node.Pool {
			conn.conn.Close()
		}
		return
	}

	if node.getSuccessor().Id != node.Id && node.getPredecessor() != nil {
		_, err := node.SetPredecessorGRPC(node.getSuccessor(), node.getPredecessor())
		if err != nil {
			NewTracer("info", "LeaveNode::SetPredecessorGRPC", err.Error())
			return
		}
		_, err = node.SetSuccessorGRPC(node.getPredecessor(), node.getSuccessor())
		if err != nil {
			NewTracer("info", "LeaveNode::SetSuccessorGRPC", err.Error())
			return
		}
	}

	return
}

func (node *Node) createChordOrJoinNode(parentNode string) error {
	if parentNode == "" {
		node.create()
		node.StartAsyncProcess()
		return nil
	}
	return node.join(parentNode)
}

func (node *Node) create() {
	node.Predecessor = nil
	node.setSuccessor(node.Node)

}

func (node *Node) join(parentNode string) error {
	node.Predecessor = nil
	successor, err := node.FindSuccessorGRPC(&protos.Node{Address: parentNode}, node.Id)
	if err != nil {
		return err
	}

	node.setSuccessor(successor)

	node.StartAsyncProcess()

	return nil
}

func (node *Node) closestPrecedingNode(id int64) *protos.Node {
	for _, fingerNode := range node.FingerTable {
		if fingerNode != nil {
			if BetweenID(fingerNode.Id, node.Id, id) {
				return fingerNode
			}
		}
	}

	return node.Node
}

func (node *Node) findSuccessor(id int64) (*protos.Node, error) {
	successor := node.getSuccessor()

	if successor == nil {
		return node.Node, nil
	}

	if BetweenID(id, node.Id, successor.Id) || id == successor.Id {
		return successor, nil
	} else {
		closestNode := node.closestPrecedingNode(id)

		if closestNode.Id == node.Node.Id {
			successor, err := node.GetSuccessorGRPC(closestNode)
			if err != nil {
				return nil, err
			}

			if successor == nil {
				return closestNode, nil
			}

			return successor, nil
		}

		successor, err := node.FindSuccessorGRPC(closestNode, id)
		if err != nil {
			return nil, err
		}

		if successor == nil {
			return node.Node, nil
		}

		return successor, nil
	}
}

func (node *Node) getSuccessor() *protos.Node {
	return node.FingerTable[0]
}

func (node *Node) setSuccessor(newSuccessor *protos.Node) {
	node.FingerTableMtx.Lock()
	node.applySetNode("SUCCESSOR", newSuccessor)
	node.FingerTableMtx.Unlock()
}

func (node *Node) getPredecessor() *protos.Node {
	predecessor := node.Predecessor

	if predecessor == nil {
		return &protos.Node{}
	}

	return predecessor
}

func (node *Node) setPredecessor(newPredecessor *protos.Node) {
	node.PredecessorMtx.Lock()
	node.applySetNode("PREDECESSOR", newPredecessor)
	node.PredecessorMtx.Unlock()
}

func (node *Node) stabilize() {
	successor := node.getSuccessor()

	if successor == nil {
		return
	}

	x, _ := node.GetPredecessorGRPC(successor)

	if x == nil {
		return
	}

	if BetweenID(x.Id, node.Id, successor.Id) {
		node.setSuccessor(x)
	}

	node.NotifyGRPC(node.getSuccessor(), node.Node)
}

func (node *Node) notify(x *protos.Node) {
	if node.Predecessor == nil || BetweenID(x.Id, node.Predecessor.Id, node.Id) {
		node.setPredecessor(x)
	}
}

func (node *Node) fingerStart(i int) int64 {
	a := math.Pow(2, float64(i))
	b := math.Pow(2, float64(GetIntEnv("CHORD_SIZE")))
	sum := float64(node.Node.Id) + a
	mod := math.Mod(sum, b)

	return int64(mod)
}

func (node *Node) fixFingerTable(count int) int {
	count = (count + 1) % GetIntEnv("CHORD_SIZE")
	fingerStart := node.fingerStart(count)
	successor, err := node.findSuccessor(fingerStart)
	if err != nil {
		return 0
	}

	node.applySetFingerRow(successor, count)

	return count
}

func (node *Node) String() string {
	s := fmt.Sprintf("Address: %s - ID: %d\n", node.Address, node.Id)

	s += "id | start | successor\n"
	for i := 0; i < len(node.FingerTable); i++ {
		if node.FingerTable[i] != nil {
			s += fmt.Sprintf("%d  | %d     | %d\n", i, node.fingerStart(i), node.FingerTable[i].Id)
		}
	}

	if node.Predecessor != nil {
		s += fmt.Sprintf("Predecessor: %d\n", node.Predecessor.Id)
	} else {
		s += fmt.Sprintf("Predecessor: None\n")
	}

	s += fmt.Sprintf("---------------\n")
	s += fmt.Sprintf("Store: \n")
	for key, value := range node.Storage {
		s += fmt.Sprintf("KEY: %d - VALUE: %s\n", key, value)
	}
	s += fmt.Sprintf("---------------\n")
	s += fmt.Sprintf("Raft: \n")
	s += fmt.Sprintf("Address: %s - ID: %d\n", node.RaftAddress, node.RaftID)
	s += fmt.Sprintf("Leader: %s\n", node.Leader)
	s += fmt.Sprintf("Replicas: \n")
	for index, replica := range node.Replicas {
		s += fmt.Sprintf("%d: %s\n", index, replica.Address)
	}

	return s
}

func (node *Node) StorageGet(key int64) (string, error) {
	closestNode, err := node.findSuccessor(key)
	if err != nil {
		return "", err
	}

	if closestNode.Address == node.Address {
		NewTracer("info", "storeGet", fmt.Sprintf("Dado recuperado do node %s", closestNode.Address))
		return node.Storage[key], nil
	}

	result, err := node.StorageGetGRPC(closestNode, key)
	if err != nil {
		return "", err
	}

	return result.Value, nil
}

func (node *Node) StorageSet(key int64, value string) error {
	node.StorageMtx.Lock()
	defer node.StorageMtx.Unlock()

	closestNode, err := node.findSuccessor(key)
	if err != nil {
		return err
	}

	if closestNode.Address == node.Address {
		node.applySet(key, value)
		NewTracer("info", "storeSet", fmt.Sprintf("Dado inserido no node %s", closestNode.Address))
		return nil
	}

	err = node.StorageSetGRPC(closestNode, key, value)
	if err != nil {
		return err
	}

	return nil
}

func (node *Node) StorageDelete(key int64) error {
	node.StorageMtx.Lock()
	defer node.StorageMtx.Unlock()

	closestNode, err := node.findSuccessor(key)
	if err != nil {
		return err
	}

	if closestNode.Address == node.Address {
		node.applyDelete(key)
		NewTracer("info", "StorageDelete", fmt.Sprintf("Dado removido no node %s", closestNode.Address))
		return nil
	}

	err = node.StorageDeleteGRPC(closestNode, key)
	if err != nil {
		return err
	}

	return nil
}
