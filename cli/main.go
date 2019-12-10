package main

import (
	"context"
	"fmt"
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
			Name:    "createChord",
			Aliases: []string{"cd", "create"},
			Usage:   "Create chord (string)",
			Flags: []cli.Flag{
				&cli.BoolFlag{Name: "watch, w"},
			},
			Action: func(c *cli.Context) {
				node, err := node.NewNode(c.Args().Get(0), "", 0)
				if err != nil {
					panic(err)
				}

				if c.Bool("watch") {
					for {
						time.Sleep(1000 * time.Millisecond)
						fmt.Printf(node.String())
					}
				}
			},
		},
		{
			Name:    "joinNode",
			Aliases: []string{"j", "join"},
			Usage:   "Join node in chord - (string string)",
			Flags: []cli.Flag{
				&cli.BoolFlag{Name: "watch, w"},
			},
			Action: func(c *cli.Context) {
				node, err := node.NewNode(c.Args().Get(0), c.Args().Get(1), -1)
				if err != nil {
					panic(err)
				}

				if c.Bool("watch") {
					for {
						time.Sleep(1000 * time.Millisecond)
						fmt.Printf(node.String())
					}
				}
			},
		},
		{
			Name:    "storageGet",
			Aliases: []string{"g", "get"},
			Usage:   "Get data in node - (string int)",
			Action: func(c *cli.Context) {
				conn, err := NewGrpcConn(c.Args().First())
				if err != nil {
					fmt.Printf("Connection error: ", err.Error())
					return
				}

				key, _ := strconv.Atoi(c.Args().Get(1))
				value, _ := conn.StorageGetRPC(context.Background(), &protos.Key{Key: int64(key)})

				fmt.Printf("VALUE = %s\n", value.Value)
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
