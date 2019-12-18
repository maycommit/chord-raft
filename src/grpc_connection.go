package sdproject

import (
	"context"
	"sdproject/protos"
	"time"

	"google.golang.org/grpc"
)

type GrpcConn struct {
	addr   string
	client protos.ChordClient
	conn   *grpc.ClientConn
}

func (node *Node) NewGrpcConn(remoteConn *protos.Node) (protos.ChordClient, error) {
	node.PoolMtx.Lock()
	grpcConn, ok := node.Pool[remoteConn.Address]
	if ok {
		node.PoolMtx.Unlock()
		return grpcConn.client, nil
	}
	node.PoolMtx.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dialOptions := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTimeout(10 * time.Second),
		grpc.FailOnNonTempDialError(true),
		grpc.WithInsecure(),
	}

	conn, err := grpc.DialContext(ctx, remoteConn.Address, dialOptions...)
	if err != nil {
		panic(err)
	}

	client := protos.NewChordClient(conn)

	grpcConn = &GrpcConn{remoteConn.Address, client, conn}

	node.PoolMtx.Lock()
	node.Pool[remoteConn.Address] = grpcConn
	node.PoolMtx.Unlock()

	return client, nil
}
