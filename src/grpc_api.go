package sdproject

import (
	"context"
	chord "sdproject/protos"
)

func (node *Node) FindSuccessorGRPC(remoteConn *chord.Node, id int64) (*chord.Node, error) {
	conn, err := node.NewGrpcConn(remoteConn)
	if err != nil {
		NewLog("error", "FindSuccessorGRPC", err.Error())
		return nil, err
	}

	result, err := conn.FindSuccessorRPC(context.Background(), &chord.ID{Id: id})
	if err != nil {
		NewLog("error", "FindSuccessorRPC", err.Error())
		return nil, err
	}

	return result, nil
}

func (node *Node) GetSuccessorGRPC(remoteConn *chord.Node) (*chord.Node, error) {
	conn, err := node.NewGrpcConn(remoteConn)
	if err != nil {
		NewLog("error", "GetSuccessorGRPC", err.Error())
		return nil, err
	}

	result, err := conn.GetSuccessorRPC(context.Background(), &chord.Any{})
	if err != nil {
		NewLog("error", "GetSuccessorRPC", err.Error())
		return nil, err
	}

	return result, nil
}

func (node *Node) GetPredecessorGRPC(remoteConn *chord.Node) (*chord.Node, error) {
	conn, err := node.NewGrpcConn(remoteConn)
	if err != nil {
		NewLog("error", "GetPredecessorGRPC", err.Error())
		return nil, err
	}

	result, err := conn.GetPredecessorRPC(context.Background(), &chord.Any{})
	if err != nil {
		NewLog("error", "GetPredecessorRPC", err.Error())
		return nil, err
	}

	return result, nil
}

func (node *Node) NotifyGRPC(remoteConn *chord.Node, x *chord.Node) error {
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
