package sdproject

import (
	"encoding/json"
	"fmt"
	"io"
	"sdproject/protos"

	"github.com/hashicorp/raft"
)

const (
	commandStoreSet       = "STORE_SET"
	commandStoreDelete    = "STORE_DELETE"
	commandSetSuccessor   = "NODE_SET_SUCCESSOR"
	commandSetPredecessor = "NODE_SET_PREDECESSOR"
	commandSetFingerTable = "NODE_SET_FINGER_TABLE"
	commandAppendReplica  = "NODE_APPEND_REPLICA"
	commandAppendConn     = "NODE_APPEND_CONN"
	commandSetLeader      = "NODE_SET_LEADER"
)

type ApplyCommand struct {
	Command string
	Key     int64
	Value   string
	Node    *protos.Node
	Count   int
	Conn    *GrpcConn
}

func (node *Node) applySet(key int64, value string) {
	command := &ApplyCommand{
		Command: commandStoreSet,
		Key:     key,
		Value:   value,
	}

	byteArray, _ := json.Marshal(&command)

	node.Raft.Apply(byteArray, GetTimeEnv("RAFT_TIMEOUT"))
}

func (node *Node) applyDelete(key int64) {
	command := &ApplyCommand{
		Command: commandStoreDelete,
		Key:     key,
		Value:   "",
	}

	byteArray, _ := json.Marshal(&command)

	node.Raft.Apply(byteArray, GetTimeEnv("RAFT_TIMEOUT"))
}

func (node *Node) applySetNode(nodeType string, newNode *protos.Node) {
	command := ""

	if nodeType == "SUCCESSOR" {
		command = commandSetSuccessor
	} else if nodeType == "PREDECESSOR" {
		command = commandSetPredecessor
	} else if nodeType == "REPLICA" {
		command = commandAppendReplica
	} else if nodeType == "LEADER" {
		command = commandSetLeader
	}

	commandApply := &ApplyCommand{
		Command: command,
		Node:    newNode,
	}

	byteArray, _ := json.Marshal(&commandApply)

	node.Raft.Apply(byteArray, GetTimeEnv("RAFT_TIMEOUT"))
}

func (node *Node) applySetFingerRow(newNode *protos.Node, count int) {
	commandApply := &ApplyCommand{
		Command: commandSetFingerTable,
		Node:    newNode,
		Count:   count,
	}

	byteArray, _ := json.Marshal(&commandApply)

	node.Raft.Apply(byteArray, GetTimeEnv("RAFT_TIMEOUT"))
}

func (node *Node) applySetLeader(newNode *protos.Node, count int) {
	commandApply := &ApplyCommand{
		Command: commandSetFingerTable,
		Node:    newNode,
		Count:   count,
	}

	byteArray, _ := json.Marshal(&commandApply)

	node.Raft.Apply(byteArray, GetTimeEnv("RAFT_TIMEOUT"))
}

func (node *Node) applyAppendConn(conn *GrpcConn, address string) {
	commandApply := &ApplyCommand{
		Command: commandAppendConn,
		Value:   address,
		Conn:    conn,
	}

	byteArray, _ := json.Marshal(&commandApply)

	node.Raft.Apply(byteArray, GetTimeEnv("RAFT_TIMEOUT"))
}

func (node *Node) Apply(l *raft.Log) interface{} {
	var command ApplyCommand
	if err := json.Unmarshal(l.Data, &command); err != nil {
		panic(fmt.Sprintf("failed to unmarshal command: %s", err.Error()))
	}

	switch command.Command {
	case commandStoreSet:
		node.Storage[command.Key] = command.Value
	case commandStoreDelete:
		delete(node.Storage, command.Key)
	case commandSetSuccessor:
		node.FingerTable[0] = command.Node
	case commandSetPredecessor:
		node.Predecessor = command.Node
	case commandSetFingerTable:
		node.FingerTable[command.Count] = command.Node
	case commandAppendReplica:
		node.Replicas = append(node.Replicas, command.Node)
	case commandSetLeader:
		node.Leader = command.Node
	case commandAppendConn:
		node.Pool[command.Value] = command.Conn
	}

	return nil
}

func (node *Node) Snapshot() (raft.FSMSnapshot, error) {
	return nil, nil
}

func (node *Node) Restore(r io.ReadCloser) error {
	return nil
}

func (node *Node) Persist(sink raft.SnapshotSink) error {
	return nil
}
