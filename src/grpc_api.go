package sdproject

import (
	"context"
	"sdproject/protos"
)

func (node *Node) JoinGRPC(remoteConn *protos.Node, parentNode *protos.Node) (*protos.Any, error) {
	conn, err := node.NewGrpcConn(remoteConn)
	if err != nil {
		return nil, err
	}

	result, err := conn.JoinRPC(context.Background(), parentNode)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (node *Node) FindSuccessorGRPC(remoteConn *protos.Node, id int64) (*protos.Node, error) {
	conn, err := node.NewGrpcConn(remoteConn)
	if err != nil {
		return nil, err
	}

	result, err := conn.FindSuccessorRPC(context.Background(), &protos.ID{Id: id})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (node *Node) GetSuccessorGRPC(remoteConn *protos.Node) (*protos.Node, error) {
	conn, err := node.NewGrpcConn(remoteConn)
	if err != nil {
		return nil, err
	}

	result, err := conn.GetSuccessorRPC(context.Background(), &protos.Any{})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (node *Node) GetPredecessorGRPC(remoteConn *protos.Node) (*protos.Node, error) {
	conn, err := node.NewGrpcConn(remoteConn)
	if err != nil {
		return nil, err
	}

	result, err := conn.GetPredecessorRPC(context.Background(), &protos.Any{})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (node *Node) SetPredecessorGRPC(remoteConn *protos.Node, newPredecessor *protos.Node) (*protos.Any, error) {
	conn, err := node.NewGrpcConn(remoteConn)
	if err != nil {
		NewTracer("error", "SetPredecessorGRPC", err.Error())
		return nil, err
	}

	_, err = conn.SetPredecessorRPC(context.Background(), newPredecessor)
	if err != nil {
		NewTracer("error", "SetPredecessorRPC", err.Error())
		return nil, err
	}

	return &protos.Any{}, nil
}

func (node *Node) SetSuccessorGRPC(remoteConn *protos.Node, newSuccessor *protos.Node) (*protos.Any, error) {
	conn, err := node.NewGrpcConn(remoteConn)
	if err != nil {
		NewTracer("error", "SetSuccessorGRPC", err.Error())
		return nil, err
	}

	_, err = conn.SetPredecessorRPC(context.Background(), newSuccessor)
	if err != nil {
		NewTracer("error", "SetSuccessorRPC", err.Error())
		return nil, err
	}

	return &protos.Any{}, nil
}

func (node *Node) NotifyGRPC(remoteConn *protos.Node, x *protos.Node) error {
	conn, err := node.NewGrpcConn(remoteConn)
	if err != nil {
		NewTracer("error", "NotifyGRPC", err.Error())
		return err
	}

	_, err = conn.NotifyRPC(context.Background(), x)
	if err != nil {
		NewTracer("error", "NotifyRPC", err.Error())
		return err
	}

	return nil
}

func (node *Node) StorageGetGRPC(remoteConn *protos.Node, key int64) (*protos.Value, error) {
	conn, err := node.NewGrpcConn(remoteConn)
	if err != nil {
		NewTracer("error", "StorageGetGRPC", err.Error())
		return nil, err
	}

	data, err := conn.StorageGetRPC(context.Background(), &protos.Key{Key: key})
	if err != nil {
		NewTracer("error", "StorageGetGRPC", err.Error())
		return nil, err
	}

	return data, nil
}

func (node *Node) StorageSetGRPC(remoteConn *protos.Node, key int64, value string) error {
	conn, err := node.NewGrpcConn(remoteConn)
	if err != nil {
		NewTracer("error", "StorageSetGRPC", err.Error())
		return err
	}

	_, err = conn.StorageSetRPC(context.Background(), &protos.Data{Key: key, Value: value})
	if err != nil {
		NewTracer("error", "StorageSetGRPC", err.Error())
		return err
	}

	return nil
}

func (node *Node) StorageDeleteGRPC(remoteConn *protos.Node, key int64) error {
	conn, err := node.NewGrpcConn(remoteConn)
	if err != nil {
		NewTracer("error", "StorageDeleteGRPC", err.Error())
		return err
	}

	_, err = conn.StorageDeleteRPC(context.Background(), &protos.Key{Key: key})
	if err != nil {
		NewTracer("error", "StorageDeleteGRPC", err.Error())
		return err
	}

	return nil
}

func (node *Node) StorageGetAllGRPC(remoteConn *protos.Node) (*protos.Datas, error) {
	conn, err := node.NewGrpcConn(remoteConn)
	if err != nil {
		NewTracer("error", "StorageGetAllGRPC", err.Error())
		return nil, err
	}

	datas, err := conn.StorageGetAllRPC(context.Background(), &protos.Any{})
	if err != nil {
		NewTracer("error", "StorageGetAllGRPC", err.Error())
		return nil, err
	}

	return datas, nil
}

func (node *Node) JoinRaftGRPC(remoteConn *protos.Node, multipleNodes *protos.MultipleNodes) (*protos.Node, error) {
	conn, err := node.NewGrpcConn(remoteConn)
	if err != nil {
		NewTracer("error", "JoinRaftGRPC", err.Error())
		return nil, err
	}

	leader, err := conn.JoinRaftRPC(context.Background(), multipleNodes)
	if err != nil {
		NewTracer("error", "JoinRaftRPC", err.Error())
		return nil, err
	}

	return leader, nil
}
