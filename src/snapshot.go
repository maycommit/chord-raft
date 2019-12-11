package sdproject

import (
	"fmt"
	"strings"
)

type Snapshot struct {
	Path     string
	Filename string
}

func NewSnapshot(address, snapshotPath string) *Snapshot {
	snapshot := &Snapshot{}
	snapshot.Path = snapshot.NewNodeSnapshotDir(address, snapshotPath)
	snapshot.Filename = ""

	return snapshot
}

func (snapshot *Snapshot) NewNodeSnapshotDir(address, snapshotPath string) string {
	splitAddress := strings.Split(address, ":")
	fileName := snapshotPath + "/" + IsLocalhost(splitAddress[0]) + "-" + splitAddress[1]
	_ = NewDir(fileName, "")

	return fileName
}

func (snapshot *Snapshot) NewSnapshotFile(storageData map[int64]string) string {
	fileName := snapshot.Path + "/" + BuildUniqueFileName() + ".snap"
	file, _ := NewFile(fileName)

	for key, value := range storageData {
		file.Write([]byte(fmt.Sprint(key) + " " + value + "\n"))
	}

	return fileName
}
