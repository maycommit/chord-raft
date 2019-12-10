package sdproject

import (
	"context"
	"sdproject/protos"
)

func (node *Node) FindSuccessorRPC(ctx context.Context, id *protos.ID) (*protos.Node, error) {
	successor, err := node.findSuccessor(id.Id)
	if err != nil {
		return nil, err
	}

	return successor, nil
}

func (node *Node) GetSuccessorRPC(ctx context.Context, any *protos.Any) (*protos.Node, error) {
	successor := node.getSuccessor()
	return successor, nil
}

func (node *Node) GetPredecessorRPC(ctx context.Context, any *protos.Any) (*protos.Node, error) {
	predecessor := node.getPredecessor()
	return predecessor, nil
}

func (node *Node) NotifyRPC(ctx context.Context, x *protos.Node) (*protos.Any, error) {
	node.notify(x)
	return &protos.Any{}, nil
}

func (node *Node) StorageGetRPC(ctx context.Context, key *protos.Key) (*protos.Value, error) {
	value, err := node.storeGet(key.Key)
	return &protos.Value{Value: value}, err
}

func (node *Node) StorageSetRPC(ctx context.Context, data *protos.Data) (*protos.Any, error) {
	node.storeSet(data.Key, data.Value)
	return &protos.Any{}, nil
}
