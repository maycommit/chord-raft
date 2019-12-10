package sdproject

import (
	"log"
	"testing"
	"time"

	"gotest.tools/assert"
)

func TestNodeFindSuccessor(t *testing.T) {
	node01, err := NewNode(":8001", "", -1)
	if err != nil {
		log.Println(err)
		return
	}

	node02, err := NewNode(":8002", ":8001", 4)
	if err != nil {
		log.Println(err)
		return
	}

	node03, err := NewNode(":8003", ":8002", 2)
	if err != nil {
		log.Println(err)
		return
	}

	time.Sleep(1 * time.Second)

	assert.Equal(t, node01.FingerTable[0].Id, int64(2))
	assert.Equal(t, node02.FingerTable[0].Id, int64(0))
	assert.Equal(t, node03.FingerTable[0].Id, int64(4))
}

func TestGetDataInNode(t *testing.T) {
	node01, err := NewNode(":8081", "", -1)
	if err != nil {
		log.Println(err)
		return
	}

	node02, err := NewNode(":8082", ":8081", 4)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = NewNode(":8083", ":8082", 2)
	if err != nil {
		log.Println(err)
		return
	}

	time.Sleep(1 * time.Second)

	node01.StorageSetGRPC(node02.Node, 1, "Maycon")

	result, err := node01.StorageGetGRPC(node02.Node, 1)
	if err != nil {
		log.Println(err)
		return
	}

	assert.Equal(t, result.Value, "Maycon")
}
