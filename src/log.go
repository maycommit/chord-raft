package sdproject

import (
	"fmt"
	"os"
	"strings"
)

type Log struct {
	Path     string
	Filename string
}

func NewLog(address, logsPath string) *Log {
	log := &Log{}
	log.Path = log.NewLogNodeDir(address, logsPath)
	log.Filename = log.NewBatchLogFile()

	return log
}

func (log *Log) NewLogNodeDir(address, logsPath string) string {
	splitAddress := strings.Split(address, ":")
	fileName := logsPath + "/" + IsLocalhost(splitAddress[0]) + "-" + splitAddress[1]
	_ = NewDir(fileName, "")

	return fileName
}

func (log *Log) NewBatchLogFile() string {
	fileName := log.Path + "/" + BuildUniqueFileName() + ".log"
	_, _ = NewFile(fileName)

	return fileName
}

func (log *Log) NewLogLine(data string) error {
	NewTracer("info", "NewLogLine", fmt.Sprintf("Novo dado no log %s\n", log.Filename))
	file, err := os.OpenFile(log.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		NewTracer("error", "NewLogLine::OpenFile", err.Error())
		return err
	}

	defer file.Close()

	if _, err := file.Write([]byte(data + "\n")); err != nil {
		NewTracer("error", "NewLogLine::Write", err.Error())
		return err
	}

	return nil
}

func (log *Log) CreateLogData(key int64, value string) string {
	return "SET " + fmt.Sprint(key) + " " + value
}
