package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"sdproject/protos"
	node "sdproject/src"

	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

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

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "readFile",
			Aliases: []string{"r", "read"},
			Usage:   "Read node file - {filePath(srtring)}",
			Action: func(c *cli.Context) {
				nodeJSON, err := os.Open(c.Args().Get(0))
				if err != nil {
					panic(err)
				}

				byteValue, _ := ioutil.ReadAll(nodeJSON)
				var nodeData *node.NodeData
				json.Unmarshal([]byte(byteValue), &nodeData)

				node, err := node.NewNode(nodeData)
				if err != nil {
					panic(err)
				}

				for {
					time.Sleep(2000 * time.Millisecond)
					fmt.Printf(node.String())
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
