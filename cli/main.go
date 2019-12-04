package main

import (
	"fmt"
	"log"
	"os"
	"time"

	node "sdproject/src"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

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
			Aliases: []string{"cd"},
			Usage:   "Create chord (First node)",
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
			Aliases: []string{"j"},
			Usage:   "Join node in chord - (newNode, chordNode)",
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
