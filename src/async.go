package sdproject

import (
	"sdproject/protos"
	"time"

	"github.com/hashicorp/raft"
)

func (node *Node) asyncFixStorage() {
	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			node.fixStorage()
		case <-node.StopNode:
			ticker.Stop()
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
			ticker.Stop()
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
			ticker.Stop()
			return
		}
	}
}

func (node *Node) asyncFixRaftLeader() {
	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			if node.Raft.State() == raft.Leader {
				node.applySetNode("LEADER", &protos.Node{Id: node.Id, Address: node.Address})
			}
		case <-node.StopNode:
			ticker.Stop()
			return
		}
	}
}
