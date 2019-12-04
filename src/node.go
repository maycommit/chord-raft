package sdproject

import (
	"fmt"
	"math"
	"math/rand"
	"net"
	chord "sdproject/protos"
	"time"

	"google.golang.org/grpc"
)

type Node struct {
	*chord.Node
	chord.ChordServer
	Predecessor *chord.Node
	Storage     *Storage
	Config      *Config
	FingerTable []*chord.Node
	Pool        map[string]*GrpcConn
}

func NewNode(address, parentNode string, id int64) (*Node, error) {
	node := &Node{
		Node:   new(chord.Node),
		Config: NewConfig(),
	}
	node.Id = node.newId(parentNode, id)
	node.Address = address
	node.Pool = make(map[string]*GrpcConn)
	node.Storage = NewStorage()
	node.FingerTable = make([]*chord.Node, node.Config.ChordSize)

	listen, err := node.startTCPServer()
	if err != nil {
		return nil, err
	}

	grpcServer := grpc.NewServer()
	chord.RegisterChordServer(grpcServer, node)

	node.createChordOrJoinNode(parentNode)

	go grpcServer.Serve(listen)

	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		for {
			select {
			case <-ticker.C:
				node.stabilize()
			}
		}
	}()

	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		next := 0
		for {
			select {
			case <-ticker.C:
				next = node.fixFingerTable(next)
			}
		}
	}()

	return node, nil
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
	max := math.Pow(2, float64(node.Config.ChordSize))

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
	successor, err := node.FindSuccessorGRPC(&chord.Node{Address: parentNode}, node.Id)
	if err != nil {
		return err
	}

	node.setSuccessor(successor)

	return nil
}

func (node *Node) closestPrecedingNode(id int64) *chord.Node {
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

func (node *Node) findSuccessor(id int64) (*chord.Node, error) {
	successor := node.getSuccessor()

	if successor == nil {
		return node.Node, nil
	}

	if node.betweenID(id, node.Id, successor.Id) {
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

func (node *Node) getSuccessor() *chord.Node {
	return node.FingerTable[0]
}

func (node *Node) setSuccessor(newSuccessor *chord.Node) {
	node.FingerTable[0] = newSuccessor
}

func (node *Node) getPredecessor() *chord.Node {
	predecessor := node.Predecessor

	if predecessor == nil {
		return &chord.Node{}
	}

	return predecessor
}

func (node *Node) setPredecessor(newPredecessor *chord.Node) {
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

func (node *Node) notify(x *chord.Node) {
	predecessor := node.Predecessor
	if predecessor == nil || node.betweenID(x.Id, predecessor.Id, node.Id) {
		node.setPredecessor(x)
	}
}

func (node *Node) fingerStart(i int) int64 {
	a := math.Pow(2, float64(i))
	b := math.Pow(2, float64(node.Config.ChordSize))
	sum := float64(node.Node.Id) + a
	mod := math.Mod(sum, b)

	return int64(mod)
}

func (node *Node) fixFingerTable(count int) int {
	count = (count + 1) % node.Config.ChordSize
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
