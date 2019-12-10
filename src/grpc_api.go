package sdproject

import (
	"context"
	"sdproject/protos"
)

func (node *Node) FindSuccessorGRPC(remoteConn *protos.Node, id int64) (*protos.Node, error) {
	conn, err := node.NewGrpcConn(remoteConn)
	if err != nil {
		NewLog("error", "FindSuccessorGRPC", err.Error())
		return nil, err
	}

	result, err := conn.FindSuccessorRPC(context.Background(), &protos.ID{Id: id})
	if err != nil {
		NewLog("error", "FindSuccessorRPC", err.Error())
		return nil, err
	}

	return result, nil
}

func (node *Node) GetSuccessorGRPC(remoteConn *protos.Node) (*protos.Node, error) {
	conn, err := node.NewGrpcConn(remoteConn)
	if err != nil {
		NewLog("error", "GetSuccessorGRPC", err.Error())
		return nil, err
	}

	result, err := conn.GetSuccessorRPC(context.Background(), &protos.Any{})
	if err != nil {
		NewLog("error", "GetSuccessorRPC", err.Error())
		return nil, err
	}

	return result, nil
}

func (node *Node) GetPredecessorGRPC(remoteConn *protos.Node) (*protos.Node, error) {
	conn, err := node.NewGrpcConn(remoteConn)
	if err != nil {
		NewLog("error", "GetPredecessorGRPC", err.Error())
		return nil, err
	}

	result, err := conn.GetPredecessorRPC(context.Background(), &protos.Any{})
	if err != nil {
		NewLog("error", "GetPredecessorRPC", err.Error())
		return nil, err
	}

	return result, nil
}

func (node *Node) NotifyGRPC(remoteConn *protos.Node, x *protos.Node) error {
	conn, err := node.NewGrpcConn(remoteConn)
	if err != nil {
		NewLog("error", "NotifyGRPC", err.Error())
		return err
	}

	_, err = conn.NotifyRPC(context.Background(), x)
	if err != nil {
		NewLog("error", "NotifyRPC", err.Error())
		return err
	}

	return nil
}

func (node *Node) StorageGetGRPC(remoteConn *protos.Node, key int64) (*protos.Value, error) {
	conn, err := node.NewGrpcConn(remoteConn)
	if err != nil {
		NewLog("error", "StorageGetGRPC", err.Error())
		return nil, err
	}

	data, err := conn.StorageGetRPC(context.Background(), &protos.Key{Key: key})
	if err != nil {
		NewLog("error", "StorageGetGRPC", err.Error())
		return nil, err
	}

	return data, nil
}

func (node *Node) StorageSetGRPC(remoteConn *protos.Node, key int64, value string) error {
	conn, err := node.NewGrpcConn(remoteConn)
	if err != nil {
		NewLog("error", "StorageSetGRPC", err.Error())
		return err
	}

	_, err = conn.StorageSetRPC(context.Background(), &protos.Data{Key: key, Value: value})
	if err != nil {
		NewLog("error", "StorageSetGRPC", err.Error())
		return err
	}

	return nil
}

func (node *Node) StorageGetAllGRPC(remoteConn *protos.Node) (*protos.Datas, error) {
	conn, err := node.NewGrpcConn(remoteConn)
	if err != nil {
		NewLog("error", "StorageGetAllGRPC", err.Error())
		return nil, err
	}

	datas, err := conn.StorageGetAllRPC(context.Background(), &protos.Any{})
	if err != nil {
		NewLog("error", "StorageGetAllGRPC", err.Error())
		return nil, err
	}

	return datas, nil
}
