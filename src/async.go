package sdproject

import "time"

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
