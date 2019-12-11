package sdproject

import (
	"os"
	"time"
)

func NewDir(dirName, path string) error {
	dirPath := path + dirName
	err := os.MkdirAll(dirPath, os.FileMode(0777))

	if err == nil || os.IsExist(err) {
		return nil
	}

	return err
}

func BuildUniqueFileName() string {
	return time.Now().Format("20060102150405")
}

func IsLocalhost(host string) string {
	if host == "" {
		return "localhost"
	}

	return host
}

func NewFile(fileName string) (*os.File, error) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}

	return file, nil
}
