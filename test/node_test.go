package sdproject

import (
	"log"
	"testing"
	"time"

	node "sdproject/src"

	"gotest.tools/assert"
)

func TestNodeFindSuccessor(t *testing.T) {
	node01, err := node.NewNode(":8001", "", -1)
	if err != nil {
		log.Println(err)
		return
	}

	node02, err := node.NewNode(":8002", ":8001", 4)
	if err != nil {
		log.Println(err)
		return
	}

	node03, err := node.NewNode(":8003", ":8002", 2)
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
	node01, err := node.NewNode(":8081", "", -1)
	if err != nil {
		log.Println(err)
		return
	}

	node02, err := node.NewNode(":8082", ":8081", 4)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = node.NewNode(":8083", ":8082", 2)
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

func TestTransferKeys(t *testing.T) {
	node01, err := node.NewNode(":8091", "", -1)
	if err != nil {
		log.Println(err)
		return
	}

	node01.StorageSet(1, "Maycon")
	node01.StorageSet(2, "Lucas")
	node01.StorageSet(4, "Antonio")
	node01.StorageSet(6, "Machado")

	node02, err := node.NewNode(":8092", ":8091", 4)
	if err != nil {
		log.Println(err)
		return
	}

	node03, err := node.NewNode(":8093", ":8092", 2)
	if err != nil {
		log.Println(err)
		return
	}
	time.Sleep(1 * time.Second)

	storageDataNode02, _ := node02.Storage.Get(4)
	storageDataNode03First, _ := node03.Storage.Get(1)
	storageDataNode03Second, _ := node03.Storage.Get(2)

	assert.Equal(t, storageDataNode02, "Antonio")
	assert.Equal(t, storageDataNode03First, "Maycon")
	assert.Equal(t, storageDataNode03Second, "Lucas")
}
