package sdproject

import (
	"fmt"
	"math"
	"math/rand"
	"net"
	"sdproject/protos"
	"time"

	"google.golang.org/grpc"
)

type Node struct {
	*protos.Node
	protos.ChordServer
	Predecessor *protos.Node
	Storage     *Storage
	FingerTable []*protos.Node
	Pool        map[string]*GrpcConn
	Log         *Log
	Snapshot    *Snapshot
	StopNode    chan struct{}
}

func NewNode(address, parentNode string, id int64) (*Node, error) {
	NewConfig()
	node := &Node{}
	node.Node = new(protos.Node)
	node.Id = node.newId(parentNode, id)
	node.Address = address
	node.StopNode = make(chan struct{})
	node.Pool = make(map[string]*GrpcConn)
	node.FingerTable = make([]*protos.Node, GetIntEnv("CHORD_SIZE"))
	node.Log = NewLog(address, GetEnv("LOGS_PATH"))
	node.Snapshot = NewSnapshot(address, GetEnv("SNAPSHOTS_PATH"), node.Log.Path)
	node.Storage = NewStorage(node.Log, node.Snapshot.GetLatestSnapshotData())

	listen, err := node.startTCPServer()
	if err != nil {
		return nil, err
	}

	grpcServer := grpc.NewServer()
	protos.RegisterChordServer(grpcServer, node)

	node.createChordOrJoinNode(parentNode)

	go grpcServer.Serve(listen)
	go node.asyncStabilize()
	go node.asyncFixFingerTable()
	go node.asyncFlushMemory()
	go node.asyncFixStorage()

	return node, nil
}

func (node *Node) asyncFixStorage() {
	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			node.fixStorage()
		case <-node.StopNode:
			node.LeaveNode()
			return
		}
	}
}

func (node *Node) asyncStabilize() {
	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			node.stabilize()
		case <-node.StopNode:
			node.LeaveNode()
			return
		}
	}
}

func (node *Node) asyncFixFingerTable() {
	ticker := time.NewTicker(100 * time.Millisecond)
	next := 0
	for {
		select {
		case <-ticker.C:
			next = node.fixFingerTable(next)
		case <-node.StopNode:
			node.LeaveNode()
			return
		}
	}
}

func (node *Node) asyncFlushMemory() {
	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			node.flushMemory()
		case <-node.StopNode:
			node.LeaveNode()
			return
		}
	}
}

func (node *Node) fixStorage() {
	for key, value := range node.Storage.Data {
		closestNode, _ := node.findSuccessor(key)

		if closestNode.Id == node.Id {
			return
		}

		node.StorageSetGRPC(closestNode, key, value)
		node.Storage.Delete(key)
	}
}

func (node *Node) LeaveNode() {
	close(node.StopNode)
	return
}

func (node *Node) startTCPServer() (net.Listener, error) {
	listen, err := net.Listen("tcp", node.Address)
	if err != nil {
		return nil, err
	}

	return listen, nil
}

func (node *Node) newId(parentNode string, id int64) int64 {
	min := 1
	max := math.Pow(2, float64(GetIntEnv("CHORD_SIZE")))

	if parentNode == "" {
		return 0
	}

	if id > 0 {
		return id
	}

	rand.Seed(time.Now().UnixNano())
	return int64(rand.Intn(int(max)-min) + min)
}

func (node *Node) createChordOrJoinNode(parentNode string) error {
	if parentNode == "" {
		node.create()
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

	return nil
}

func (node *Node) closestPrecedingNode(id int64) *protos.Node {
	for _, fingerNode := range node.FingerTable {
		if fingerNode != nil {
			if node.betweenID(fingerNode.Id, node.Id, id) {
				return fingerNode
			}
		}
	}

	return node.Node
}

func (node *Node) betweenID(id, init, end int64) bool {
	if init < end {
		return id > init && id < end
	}

	if init > end {
		return id > init || id < end
	}

	if init == end {
		return id > init || id < init
	}

	return false
}

func (node *Node) findSuccessor(id int64) (*protos.Node, error) {
	successor := node.getSuccessor()

	if successor == nil {
		return node.Node, nil
	}

	if node.betweenID(id, node.Id, successor.Id) || id == successor.Id {
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
	node.FingerTable[0] = newSuccessor
}

func (node *Node) getPredecessor() *protos.Node {
	predecessor := node.Predecessor

	if predecessor == nil {
		return &protos.Node{}
	}

	return predecessor
}

func (node *Node) setPredecessor(newPredecessor *protos.Node) {
	node.Predecessor = newPredecessor
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

	if node.betweenID(x.Id, node.Id, successor.Id) {
		node.setSuccessor(x)
	}

	node.NotifyGRPC(node.getSuccessor(), node.Node)
}

func (node *Node) transferKeys(target *protos.Node, init int64, end int64) {
	count := 0
	for key, value := range node.Storage.Data {
		if node.betweenID(key, init, end) || key == end {
			node.StorageSetGRPC(target, key, value)
			delete(node.Storage.Data, key)
			count++
		}
	}
}

func (node *Node) notify(x *protos.Node) {
	if node.Predecessor == nil || node.betweenID(x.Id, node.Predecessor.Id, node.Id) {
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
		return node.Storage.Get(key)
	}

	result, err := node.StorageGetGRPC(closestNode, key)
	if err != nil {
		return "", err
	}

	return result.Value, nil
}

func (node *Node) StorageSet(key int64, value string) error {
	closestNode, err := node.findSuccessor(key)
	if err != nil {
		return err
	}

	if closestNode.Address == node.Address {
		node.Storage.Set(key, value)
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
	closestNode, err := node.findSuccessor(key)
	if err != nil {
		return err
	}

	if closestNode.Address == node.Address {
		node.Storage.Delete(key)
		NewTracer("info", "StorageDelete", fmt.Sprintf("Dado removido no node %s", closestNode.Address))
		return nil
	}

	err = node.StorageDeleteGRPC(closestNode, key)
	if err != nil {
		return err
	}

	return nil
}

func (node *Node) flushMemory() {
	if !GetBoolEnv("PERSISTENCE") {
		return
	}

	if node.Storage.SnapshotTrigger < GetIntEnv("SNAPSHOT_MAX_TRIGGER") {
		return
	}

	node.Snapshot.NewSnapshotFile(node.Storage.Data)
	node.Storage.SnapshotTriggerClear()
}
