package main

import (
	"fmt"
	"log"
	node "sdproject/src"
	"time"
)

func main() {
	node01, err := node.NewNode(":8001", "", -1)
	node02, err := node.NewNode(":8002", ":8001", 4)
	node03, err := node.NewNode(":8004", ":8002", 2)
	if err != nil {
		log.Println("Erro ao iniciar no: ", err.Error())
		return
	}

	for {
		time.Sleep(5 * time.Second)
		fmt.Println(node01.String())
		fmt.Println(node02.String())
		fmt.Println(node03.String())
	}
}
