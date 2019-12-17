package sdproject

import (
	"encoding/json"
	"io"

	"github.com/hashicorp/raft"
)

type StoreApplyCommand struct {
	Command string
	Key     int64
	Value   string
}

func (node *Node) applySet(key int64, value string) {
	command := &StoreApplyCommand{
		Command: "SET",
		Key:     key,
		Value:   value,
	}

	byteArray, _ := json.Marshal(&command)

	node.Raft.Apply(byteArray, GetTimeEnv("RAFT_TIMEOUT"))
}

func (node *Node) applyDelete(key int64) {
	command := &StoreApplyCommand{
		Command: "DELETE",
		Key:     key,
		Value:   "",
	}

	byteArray, _ := json.Marshal(&command)

	node.Raft.Apply(byteArray, GetTimeEnv("RAFT_TIMEOUT"))
}

func (node *Node) Apply(l *raft.Log) interface{} {
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
