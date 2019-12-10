package sdproject

import (
	"errors"
)

type Storage struct {
	Data map[int64]string
}

func NewStorage() *Storage {
	return &Storage{
		Data: make(map[int64]string),
	}
}

func (storage *Storage) Get(key int64) (string, error) {
	value, ok := storage.Data[key]
	if !ok {
		return "", errors.New("Data not found")
	}

	return value, nil
}

func (storage *Storage) Set(key int64, value string) {
	storage.Data[key] = value
}
