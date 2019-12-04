package sdproject

import (
	"log"
	"testing"
	"time"

	"gotest.tools/assert"
)

func TestNodeFindSuccessor(t *testing.T) {
	node01, err := NewNode(":8001", "", -1)
	node02, err := NewNode(":8002", ":8001", 4)
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
