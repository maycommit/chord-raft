package sdproject

import (
	"context"
	chord "sdproject/protos"
	"time"

	"google.golang.org/grpc"
)

type GrpcConn struct {
	addr   string
	client chord.ChordClient
	conn   *grpc.ClientConn
}

func (node *Node) NewGrpcConn(remoteConn *chord.Node) (chord.ChordClient, error) {
	grpcConn, ok := node.Pool[remoteConn.Address]
	if ok {
		return grpcConn.client, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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

	client := chord.NewChordClient(conn)

	grpcConn = &GrpcConn{remoteConn.Address, client, conn}
	node.Pool[remoteConn.Address] = grpcConn

	return client, nil
}
