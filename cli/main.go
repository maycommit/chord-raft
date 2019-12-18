package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	"sdproject/protos"
	node "sdproject/src"

	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

type NodeFileData struct {
	*node.NodeData
	Replicas []*node.NodeData
}

var app = cli.NewApp()

func NewGrpcConn(remoteConn string) (protos.ChordClient, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dialOptions := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTimeout(10 * time.Second),
		grpc.FailOnNonTempDialError(true),
		grpc.WithInsecure(),
	}

	conn, err := grpc.DialContext(ctx, remoteConn, dialOptions...)
	if err != nil {
		panic(err)
	}

	client := protos.NewChordClient(conn)

	return client, nil
}

func info() {
	app.Name = "Project CLI"
	app.Usage = "CLI for node create and client connection"
	app.Author = "Maycon Pacheco"
	app.Version = "1.0.0"
}

func transformFileDataToNodeData(nodeFileData *NodeFileData) *node.NodeData {
	return &node.NodeData{
		Id:            nodeFileData.Id,
		Address:       nodeFileData.Address,
		Parent:        nodeFileData.Parent,
		RaftAddress:   nodeFileData.RaftAddress,
		RaftDirectory: nodeFileData.RaftDirectory,
	}
}

func handleLeaveNode(node *node.Node) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		<-signalChan
		fmt.Println("\nReceived an interrupt, stopping services...\n")
		node.LeaveNode()
		close(node.StopNode)
		os.Exit(1)
	}()
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "createChord",
			Aliases: []string{"c", "create"},
			Usage:   "Create chord - {nodeAddress(string) raftId(int) raftAddress(string) raftDirectory(string)}",
			Action: func(c *cli.Context) {
				raftID, _ := strconv.Atoi(c.Args().Get(1))
				nodeData := &node.NodeData{
					Id:            0,
					Address:       c.Args().Get(0),
					RaftID:        raftID,
					RaftAddress:   c.Args().Get(2),
					RaftDirectory: c.Args().Get(3),
					Parent:        "",
				}
				newNode, err := node.NewNode(nodeData, false)
				if err != nil {
					panic(err)
				}

				handleLeaveNode(newNode)

				for {
					time.Sleep(2000 * time.Millisecond)
					fmt.Printf(newNode.String())
				}
			},
		},
		{
			Name:    "joinNode",
			Aliases: []string{"j", "join"},
			Usage:   "Join node in chord - {nodeID(int) raftID(int) nodeAddress(string) parentNode(string) raftAddress(string) raftDirectory(string)}",
			Action: func(c *cli.Context) {
				id, _ := strconv.Atoi(c.Args().Get(0))
				raftID, _ := strconv.Atoi(c.Args().Get(1))
				nodeData := &node.NodeData{
					Id:            id,
					RaftID:        raftID,
					Address:       c.Args().Get(2),
					Parent:        c.Args().Get(3),
					RaftAddress:   c.Args().Get(4),
					RaftDirectory: c.Args().Get(5),
				}
				newNode, err := node.NewNode(nodeData, false)
				if err != nil {
					panic(err)
				}

				handleLeaveNode(newNode)

				for {
					time.Sleep(2000 * time.Millisecond)
					fmt.Printf(newNode.String())
				}
			},
		},
		{
			Name:    "addReplica",
			Aliases: []string{"r", "replica"},
			Usage:   "Create node replica - {nodeID(int) raftID(int) replicaAddress(string) nodeAddress(string) raftAddress(string) raftDirectory(string)}",
			Action: func(c *cli.Context) {
				id, _ := strconv.Atoi(c.Args().Get(0))
				raftID, _ := strconv.Atoi(c.Args().Get(1))
				nodeData := &node.NodeData{
					Id:            id,
					RaftID:        raftID,
					Address:       c.Args().Get(2),
					Parent:        c.Args().Get(3),
					RaftAddress:   c.Args().Get(4),
					RaftDirectory: c.Args().Get(5),
				}
				newNode, err := node.NewNode(nodeData, true)
				if err != nil {
					panic(err)
				}

				for {
					time.Sleep(2000 * time.Millisecond)
					fmt.Printf(newNode.String())
				}
			},
		},
		{
			Name:    "storageGet",
			Aliases: []string{"g", "get"},
			Usage:   "Get data in node - {nodeAddress(string) key(int)}",
			Action: func(c *cli.Context) {
				conn, err := NewGrpcConn(c.Args().First())
				if err != nil {
					fmt.Printf("Connection error: ", err.Error())
					return
				}

				key, _ := strconv.Atoi(c.Args().Get(1))
				value, _ := conn.StorageGetRPC(context.Background(), &protos.Key{Key: int64(key)})

				fmt.Println(value)
			},
		},
		{
			Name:    "storageSet",
			Aliases: []string{"s", "set"},
			Usage:   "Set data in node - {nodeAddress(string) key(int) value(string)}",
			Action: func(c *cli.Context) {
				conn, err := NewGrpcConn(c.Args().First())
				if err != nil {
					fmt.Printf("Connection error: ", err.Error())
					return
				}

				key, _ := strconv.Atoi(c.Args().Get(1))
				value := c.Args().Get(2)
				_, err = conn.StorageSetRPC(context.Background(), &protos.Data{Key: int64(key), Value: value})
				if err != nil {
					fmt.Printf("Erro ao inserir dado!\n")
					return
				}

				fmt.Printf("Dados inserido com sucesso\n")
			},
		},
		{
			Name:    "storageDelete",
			Aliases: []string{"d", "delete"},
			Usage:   "Delete data in node - {nodeAddress(string) key(int)}",
			Action: func(c *cli.Context) {
				conn, err := NewGrpcConn(c.Args().First())
				if err != nil {
					fmt.Printf("Connection error: ", err.Error())
					return
				}

				key, _ := strconv.Atoi(c.Args().Get(1))
				_, err = conn.StorageDeleteRPC(context.Background(), &protos.Key{Key: int64(key)})
				if err != nil {
					fmt.Printf("Erro ao deletar dado!\n")
					return
				}

				fmt.Printf("Dados deletado com sucesso\n")
			},
		},
	}
}

func main() {
	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
