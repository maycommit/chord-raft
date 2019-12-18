package test

import sdproject "sdproject/src"

import "fmt"

import "strconv"

func GenerateNode(id int, address, parent, raftAddress string) *sdproject.NodeData {
	data := &sdproject.NodeData{}
	data.Id = id
	data.Address = address
	data.RaftID = id + 10
	data.Parent = parent
	data.RaftAddress = raftAddress
	data.RaftDirectory = fmt.Sprintf("./node-%s", strconv.Itoa(data.RaftID))

	return data
}

func GenerateReplica(id, raftID int, address, parent, raftAddress string) *sdproject.NodeData {
	data := &sdproject.NodeData{}
	data.Id = id
	data.Address = address
	data.RaftID = raftID
	data.Parent = parent
	data.RaftAddress = raftAddress
	data.RaftDirectory = fmt.Sprintf("./node-%s", strconv.Itoa(data.RaftID))

	return data
}
