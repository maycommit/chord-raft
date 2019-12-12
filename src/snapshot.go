package sdproject

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Snapshot struct {
	Path     string
	Filename string
}

func NewSnapshot(address, snapshotPath, logsPath string) *Snapshot {
	snapshot := &Snapshot{}
	snapshot.Path = snapshot.NewNodeSnapshotDir(address, snapshotPath)
	snapshot.Filename, _ = snapshot.LoadInitialSnapshot(logsPath)

	return snapshot
}

func (snapshot *Snapshot) getAllLogs(path string) []string {
	var logs []string
	logFiles, err := ioutil.ReadDir(path)
	if err != nil {
		return nil
	}

	for _, logFile := range logFiles {
		data, _ := ioutil.ReadFile(path + "/" + logFile.Name())
		for _, row := range strings.Split(string(data), "\n") {
			if row != "" {
				logs = append(logs, row)
			}
		}
	}

	return logs
}

func (snapshot *Snapshot) convertToStorageData(bodyFile string) map[int64]string {
	data := map[int64]string{}

	for _, row := range strings.Split(bodyFile, "\n") {
		if row != "" {
			splitRow := strings.Split(row, " ")
			key, _ := strconv.Atoi(splitRow[0])
			data[int64(key)] = splitRow[1]
		}
	}

	return data
}

func (snapshot *Snapshot) GetLatestSnapshotData() map[int64]string {
	data := map[int64]string{}

	snapshotsFiles, err := ioutil.ReadDir(snapshot.Path)
	if err != nil {
		return data
	}
	if len(snapshotsFiles) > 0 {
		latestSnapshot := snapshotsFiles[len(snapshotsFiles)-1].Name()
		bodyFile, _ := ioutil.ReadFile(snapshot.Path + "/" + latestSnapshot)
		data = snapshot.convertToStorageData(string(bodyFile))
	}

	return data
}

func (snapshot *Snapshot) LoadInitialSnapshot(logsPath string) (string, error) {
	storageData := snapshot.GetLatestSnapshotData()
	allLogs := snapshot.getAllLogs(logsPath)
	if len(allLogs) <= 0 {
		return "", nil
	}

	for _, row := range allLogs {
		splitRow := strings.Split(row, " ")
		if splitRow[0] == "SET" {
			key, _ := strconv.Atoi(splitRow[1])
			storageData[int64(key)] = splitRow[2]
		} else if splitRow[0] == "DELETE" {
			key, _ := strconv.Atoi(splitRow[1])
			delete(storageData, int64(key))
		}
	}

	fileName := snapshot.NewSnapshotFile(storageData)

	return fileName, nil
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
