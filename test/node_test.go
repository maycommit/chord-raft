package test

import (
	"fmt"
	"testing"
	"time"

	sdproject "sdproject/src"
)

func TestNodeJoin(t *testing.T) {
	fmt.Println("------- TestNodeJoin ---------")
	nodeData0 := GenerateNode(0, ":8000", "", ":12000")
	node0, err := sdproject.NewNode(nodeData0, false)
	if err != nil {
		t.Error(err.Error())
		return
	}

	nodeData1 := GenerateNode(4, ":8002", ":8000", ":12001")
	node1, err := sdproject.NewNode(nodeData1, false)
	if err != nil {
		t.Error(err.Error())
		return
	}

	nodeData2 := GenerateNode(2, ":8003", ":8000", ":12002")
	node2, err := sdproject.NewNode(nodeData2, false)
	if err != nil {
		t.Error(err.Error())
		return
	}

	time.Sleep(1 * time.Second)

	fmt.Println(node0.String())
	fmt.Println(node1.String())
	fmt.Println(node2.String())

	defer node1.LeaveNode()
}

func TestAddReplicaToNode(t *testing.T) {
	fmt.Println("------- TestAddReplicaToNode ---------")
	nodeData0 := GenerateNode(0, ":9000", "", ":13000")
	node0, err := sdproject.NewNode(nodeData0, false)
	if err != nil {
		t.Error(err.Error())
		return
	}

	replicas := []*sdproject.NodeData{
		GenerateNode(0, ":9001", ":9000", ":13001"),
		GenerateNode(0, ":9002", ":9000", ":13002"),
		GenerateNode(0, ":9003", ":9000", ":13003"),
	}

	for _, replica := range replicas {
		_, err = sdproject.NewReplica(replica)
		if err != nil {
			t.Error(err.Error())
			return
		}
	}

	time.Sleep(1 * time.Second)

	fmt.Println(node0.String())

	defer node0.LeaveNode()
}

func TestAddReplicaToChord(t *testing.T) {
	fmt.Println("------- TestAddReplicaToChord ---------")

	nodeData0 := GenerateNode(0, ":10000", "", ":14000")
	node0, err := sdproject.NewNode(nodeData0, false)
	if err != nil {
		t.Error(err.Error())
		return
	}

	nodeData1 := GenerateNode(4, ":10001", ":10000", ":14001")
	node1, err := sdproject.NewNode(nodeData1, false)
	if err != nil {
		t.Error(err.Error())
		return
	}

	replicas := []*sdproject.NodeData{
		GenerateNode(4, ":10002", ":10001", ":14002"),
		GenerateNode(4, ":10003", ":10001", ":14003"),
		GenerateNode(4, ":10004", ":10001", ":14004"),
	}

	for _, replica := range replicas {
		_, err = sdproject.NewReplica(replica)
		if err != nil {
			t.Error(err.Error())
			return
		}
	}

	time.Sleep(1 * time.Second)

	node1.LeaveNode()

	time.Sleep(1 * time.Second)

	fmt.Println(node0.String())

	defer node0.LeaveNode()
}

func TestLogReplication(t *testing.T) {
	fmt.Println("------- TestLogReplication ---------")

	nodeData0 := GenerateNode(0, ":11000", "", ":15000")
	node0, err := sdproject.NewNode(nodeData0, false)
	if err != nil {
		t.Error(err.Error())
		return
	}

	nodeData1 := GenerateNode(4, ":11001", ":11000", ":15001")
	node1, err := sdproject.NewNode(nodeData1, false)
	if err != nil {
		t.Error(err.Error())
		return
	}

	replicasData := []*sdproject.NodeData{
		GenerateReplica(4, 41, ":11002", ":11001", ":15002"),
		GenerateReplica(4, 42, ":11003", ":11001", ":15003"),
		GenerateReplica(4, 43, ":11004", ":11001", ":15004"),
	}

	replicas := []*sdproject.Node{}

	for _, replica := range replicasData {
		newReplica, err := sdproject.NewNode(replica, true)
		if err != nil {
			t.Error(err.Error())
			return
		}

		replicas = append(replicas, newReplica)
	}

	node0.StorageSet(2, "Maycon")
	node1.StorageSet(3, "Vitor")
	node0.StorageSet(4, "Joao")
	node1.StorageSet(5, "Maria")

	time.Sleep(1 * time.Second)

	fmt.Println(node0.String())
	fmt.Println(node1.String())
	fmt.Println("REPLICAS:")
	for _, replica := range replicas {
		fmt.Println(replica.String())
	}

	defer node0.LeaveNode()
}
