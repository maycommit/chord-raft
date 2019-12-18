package sdproject

import (
	"context"
	"sdproject/protos"
)

func (node *Node) JoinRPC(ctx context.Context, parentNode *protos.Node) (*protos.Any, error) {
	err := node.createChordOrJoinNode(parentNode.Address)
	if err != nil {
		return nil, err
	}

	return &protos.Any{}, nil
}

func (node *Node) LeaveRPC(ctx context.Context, leaveNode *protos.Node) (*protos.Any, error) {
	node.Pool[leaveNode.Address].conn.Close()
	delete(node.Pool, leaveNode.Address)

	return &protos.Any{}, nil
}

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

func (node *Node) SetPredecessorRPC(ctx context.Context, newPredecessor *protos.Node) (*protos.Any, error) {
	node.setPredecessor(newPredecessor)
	return &protos.Any{}, nil
}

func (node *Node) SetSuccessorRPC(ctx context.Context, newSuccessor *protos.Node) (*protos.Any, error) {
	node.setSuccessor(newSuccessor)
	return &protos.Any{}, nil
}

func (node *Node) NotifyRPC(ctx context.Context, x *protos.Node) (*protos.Any, error) {
	node.notify(x)
	return &protos.Any{}, nil
}

func (node *Node) StorageGetRPC(ctx context.Context, key *protos.Key) (*protos.Value, error) {
	value, err := node.StorageGet(key.Key)
	return &protos.Value{Value: value}, err
}

func (node *Node) StorageImediateSetRPC(ctx context.Context, data *protos.Data) (*protos.Any, error) {
	node.applySet(data.Key, data.Value)
	return &protos.Any{}, nil
}

func (node *Node) StorageImediateDeleteRPC(ctx context.Context, data *protos.Key) (*protos.Any, error) {
	node.applyDelete(data.Key)
	return &protos.Any{}, nil
}

func (node *Node) StorageSetRPC(ctx context.Context, data *protos.Data) (*protos.Any, error) {
	node.StorageSet(data.Key, data.Value)
	return &protos.Any{}, nil
}

func (node *Node) StorageDeleteRPC(ctx context.Context, data *protos.Key) (*protos.Any, error) {
	node.StorageDelete(data.Key)
	return &protos.Any{}, nil
}

func (node *Node) JoinRaftRPC(ctx context.Context, multipleNodes *protos.MultipleNodes) (*protos.Node, error) {
	leader, err := node.joinRaft(multipleNodes.ChordNode, multipleNodes.RaftNode)
	if err != nil {
		return nil, err
	}

	return leader, nil
}
