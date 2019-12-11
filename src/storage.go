package sdproject

import (
	"errors"
)

type Storage struct {
	Log             *Log
	Data            map[int64]string
	SnapshotTrigger int
}

func NewStorage(log *Log) *Storage {
	return &Storage{
		Log:             log,
		Data:            make(map[int64]string),
		SnapshotTrigger: 0,
	}
}

func (storage *Storage) SnapshotTriggerClear() {
	storage.SnapshotTrigger = 0
}

func (storage *Storage) Get(key int64) (string, error) {
	value, ok := storage.Data[key]
	if !ok {
		return "", errors.New("Data not found")
	}

	return value, nil
}

func (storage *Storage) Set(key int64, value string) {
	logData := storage.Log.CreateLogData(key, value)
	err := storage.Log.NewLogLine(logData)

	if err != nil {
		NewTracer("error", "Set", err.Error())
		return
	}

	storage.SnapshotTrigger += 1
	storage.Data[key] = value
}
