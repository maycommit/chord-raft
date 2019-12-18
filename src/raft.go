package sdproject

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"sdproject/protos"
	"time"

	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"
)

func (node *Node) raftPersistence() (raft.LogStore, raft.StableStore, error) {
	if !GetBoolEnv("PERSISTENCE") {
		logStore := raft.NewInmemStore()
		stableStore := raft.NewInmemStore()
		return logStore, stableStore, nil
	}

	os.MkdirAll(node.RaftDirectory, 0777)

	boltDB, err := raftboltdb.NewBoltStore(filepath.Join(node.RaftDirectory, "raft.db"))
	if err != nil {
		return nil, nil, err
	}

	return boltDB, boltDB, nil

}

func (node *Node) raftSnapshotStore() (raft.SnapshotStore, error) {
	if !GetBoolEnv("PERSISTENCE") {
		return raft.NewInmemSnapshotStore(), nil
	}

	snapshots, err := raft.NewFileSnapshotStore(node.RaftDirectory, GetIntEnv("SNAPSHOT_COUNT"), os.Stderr)
	if err != nil {
		return nil, fmt.Errorf("file snapshot store: %s", err)
	}

	return snapshots, nil

}

func NewReplica(data *NodeData) (*Node, error) {
	replica, err := NewNode(data, true)
	if err != nil {
		NewTracer("error", "NewReplica", err.Error())
		return nil, err
	}

	NewTracer("info", "NewReplica", fmt.Sprintf("Replica %d para o node %s criado com sucesso!", replica.Id, data.Parent))
	return replica, err
}

func (node *Node) newRaftNode(isReplica bool) (*raft.Raft, error) {
	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID(node.RaftID)
	config.LogLevel = "ERROR"

	addr, err := net.ResolveTCPAddr("tcp", node.RaftAddress)
	if err != nil {
		return nil, err
	}
	transport, err := raft.NewTCPTransport(node.RaftAddress, addr, 3, 10*time.Second, os.Stderr)
	if err != nil {
		return nil, err
	}

	logStore, stableStore, err := node.raftPersistence()
	if err != nil {
		return nil, err
	}

	snapshots, err := node.raftSnapshotStore()
	if err != nil {
		return nil, err
	}

	raftImpl, err := raft.NewRaft(config, node, logStore, stableStore, snapshots, transport)
	if err != nil {
		return nil, fmt.Errorf("new raft: %s", err)
	}

	if !isReplica {
		configuration := raft.Configuration{
			Servers: []raft.Server{
				{
					Suffrage: raft.Voter,
					ID:       config.LocalID,
					Address:  transport.LocalAddr(),
				},
			},
		}
		raftImpl.BootstrapCluster(configuration)
	}

	return raftImpl, nil
}

func (node *Node) joinRaft(chordNode, raftNode *protos.Node) (*protos.Node, error) {
	indexFuture := node.Raft.AddVoter(raft.ServerID(raftNode.Id), raft.ServerAddress(raftNode.Address), 0, 0)
	if indexFuture.Error() != nil {
		return nil, indexFuture.Error()
	}

	node.applySetNode("REPLICA", &protos.Node{Id: chordNode.Id, Address: chordNode.Address})

	leader := &protos.Node{Id: node.Id, Address: node.Address}

	return leader, nil
}
