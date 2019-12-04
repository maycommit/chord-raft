package sdproject

import (
	"context"
	chord "sdproject/protos"
)

func (node *Node) FindSuccessorRPC(ctx context.Context, id *chord.ID) (*chord.Node, error) {
	successor, err := node.findSuccessor(id.Id)
	if err != nil {
		return nil, err
	}

	return successor, nil
}

func (node *Node) GetSuccessorRPC(ctx context.Context, any *chord.Any) (*chord.Node, error) {
	successor := node.getSuccessor()
	return successor, nil
}

func (node *Node) GetPredecessorRPC(ctx context.Context, any *chord.Any) (*chord.Node, error) {
	predecessor := node.getPredecessor()
	return predecessor, nil
}

func (node *Node) NotifyRPC(ctx context.Context, x *chord.Node) (*chord.Any, error) {
	node.notify(x)
	return &chord.Any{}, nil
}
