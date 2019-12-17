package sdproject

import (
	"math"
	"math/rand"
	"net"
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

func BetweenID(id, init, end int64) bool {
	if init < end {
		return id > init && id < end
	}

	if init > end {
		return id > init || id < end
	}

	if init == end {
		return id > init || id < init
	}

	return false
}

func StartTCPServer(address string) (net.Listener, error) {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	return listen, nil
}

func NewId(parentNode string, id int64) int64 {
	min := 1
	max := math.Pow(2, float64(GetIntEnv("CHORD_SIZE")))

	if parentNode == "" {
		return 0
	}

	if id > 0 {
		return id
	}

	rand.Seed(time.Now().UnixNano())
	return int64(rand.Intn(int(max)-min) + min)
}
