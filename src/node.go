package sdproject

import (
	"fmt"
	"math"
	"net"
	"os"
	"path/filepath"
	"sdproject/protos"
	"sync"
	"time"

	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"
	"google.golang.org/grpc"
)

type Replica struct {
	Id            int
	Address       string
	RaftAddress   string
	RaftDirectory string
}

type NodeData struct {
	Replica
	Parent   string
	Replicas []Replica
}

type NodeMtx struct {
	StorageMtx     sync.RWMutex
	FingerTableMtx sync.RWMutex
	PredecessorMtx sync.RWMutex
	SnapshotMtx    sync.RWMutex
}

type Node struct {
	*protos.Node
	protos.ChordServer
	NodeMtx
	Predecessor   *protos.Node
	Storage       map[int64]string
	FingerTable   []*protos.Node
	Pool          map[string]*GrpcConn
	Replicas      []Replica
	Raft          *raft.Raft
	RaftAddress   string
	RaftDirectory string
	RaftServers   []raft.Server
	StopNode      chan struct{}
}

func NewNode(data *NodeData) (*Node, error) {
	NewConfig()
	node := &Node{}
	node.Node = new(protos.Node)
	node.Id = NewId(data.Parent, int64(data.Id))
	node.Address = data.Address
	node.RaftAddress = data.RaftAddress
	node.RaftDirectory = data.RaftDirectory
	node.Replicas = data.Replicas
	node.StopNode = make(chan struct{})
	node.Pool = make(map[string]*GrpcConn)
	node.FingerTable = make([]*protos.Node, GetIntEnv("CHORD_SIZE"))
	node.Storage = make(map[int64]string)
	node.RaftServers = node.createRaftServers(data.Replicas)

	listen, err := StartTCPServer(data.Address)
	if err != nil {
		return nil, err
	}

	node.newRaftCluster()

	grpcServer := grpc.NewServer()
	protos.RegisterChordServer(grpcServer, node)

	node.createChordOrJoinNode(data.Parent)

	go grpcServer.Serve(listen)
	go node.asyncStabilize()
	go node.asyncFixFingerTable()
	go node.asyncFixStorage()

	return node, nil
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
	close(node.StopNode)
	return
}

func (node *Node) createChordOrJoinNode(parentNode string) error {
	if parentNode == "" {
		node.create()
		return nil
	}
	return node.join(parentNode)
}

func (node *Node) raftPersistence() (raft.LogStore, raft.StableStore, error) {
	if !GetBoolEnv("PERSISTENCE") {
		logStore := raft.NewInmemStore()
		stableStore := raft.NewInmemStore()
		return logStore, stableStore, nil
	}

	os.MkdirAll(node.RaftDirectory, 0777)

	boltDB, err := raftboltdb.NewBoltStore(filepath.Join(node.RaftDirectory, "raft.db"))
	if err != nil {
		return nil, nil, err
	}

	return boltDB, boltDB, nil

}

func (node *Node) raftSnapshotStore() (raft.SnapshotStore, error) {
	if !GetBoolEnv("PERSISTENCE") {
		return raft.NewInmemSnapshotStore(), nil
	}

	snapshots, err := raft.NewFileSnapshotStore(node.RaftDirectory, GetIntEnv("SNAPSHOT_COUNT"), os.Stderr)
	if err != nil {
		return nil, fmt.Errorf("file snapshot store: %s", err)
	}

	return snapshots, nil

}

func (node *Node) newRaftCluster() {
	node.Raft, _ = node.newRaftNode(int(node.Id), node.RaftAddress, node.RaftDirectory)
	for _, replica := range node.Replicas {
		_, err := node.newRaftNode(replica.Id, replica.RaftAddress, replica.RaftDirectory)
		if err != nil {
			NewTracer("error", "newRaftCluster::newRaftNode", err.Error())
			return
		}

		listen, err := StartTCPServer(replica.Address)
		if err != nil {
			fmt.Println("Error replica start")
			return
		}
		NewTracer("info", "newRaftCluster", fmt.Sprintf("Repica %s started", replica.Address))

		grpcServer := grpc.NewServer()
		protos.RegisterChordServer(grpcServer, node)
		go grpcServer.Serve(listen)
	}

	NewTracer("info", "newRaftCluster", "Cluster created!!!")
}

func (node *Node) createRaftServers(replicas []Replica) []raft.Server {
	servers := make([]raft.Server, 0, len(replicas)+1)
	servers = append(servers, raft.Server{
		Suffrage: raft.Voter,
		ID:       raft.ServerID(node.Id),
		Address:  raft.ServerAddress(node.RaftAddress),
	})

	for _, replica := range replicas {
		servers = append(servers, raft.Server{
			Suffrage: raft.Voter,
			ID:       raft.ServerID(replica.Id),
			Address:  raft.ServerAddress(replica.RaftAddress),
		})
	}

	return servers
}

func (node *Node) newRaftNode(id int, raftAddress, raftDirectory string) (*raft.Raft, error) {
	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID(id)

	addr, err := net.ResolveTCPAddr("tcp", raftAddress)
	if err != nil {
		return nil, err
	}
	transport, err := raft.NewTCPTransport(raftAddress, addr, 3, 10*time.Second, os.Stderr)
	if err != nil {
		return nil, err
	}

	logStore, stableStore, err := node.raftPersistence()
	if err != nil {
		return nil, err
	}

	snapshots, err := node.raftSnapshotStore()
	if err != nil {
		return nil, err
	}

	configuration := raft.Configuration{Servers: node.RaftServers}

	if err = raft.BootstrapCluster(config, logStore, stableStore, snapshots, transport, configuration); err != nil {
		return nil, err
	}

	raftImpl, err := raft.NewRaft(config, node, logStore, stableStore, snapshots, transport)
	if err != nil {
		return nil, fmt.Errorf("new raft: %s", err)
	}

	return raftImpl, nil
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
	node.FingerTable[0] = newSuccessor
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
	node.Predecessor = newPredecessor
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
		return count
	}

	node.FingerTable[count] = successor

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
		s += fmt.Sprintf("Predecessor: %d", node.Predecessor.Id)
	} else {
		s += fmt.Sprintf("Predecessor: None")
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
