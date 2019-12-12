package sdproject

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const projectPath = "~/.sdproject"
const logsPath = "./logs"
const snapsPath = "./snaps"

func NewConfig() {
	configs := map[string]string{
		"CHORD_SIZE":                "3",
		"MEMORY_SIZE":               "3",
		"GRPC_TIMEOUT":              "8",
		"STABILIZE_INTERVAL":        "100",
		"FIX_FINGER_TABLE_INTERVAL": "100",
		"FLUSH_MEMORY_INTERVAL":     "100",
		"SNAPSHOT_MAX_TRIGGER":      "3",
		"PERSISTENCE":               "true",
		"LOGS_PATH":                 logsPath,
		"SNAPSHOTS_PATH":            snapsPath,
	}

	for key, value := range configs {
		os.Setenv(key, value)
	}

	fileEnv, err := godotenv.Read("./config.yaml")
	if err != nil {
		return
	}

	for key, value := range fileEnv {
		os.Setenv(key, value)
	}
}

func GetIntEnv(key string) int {
	result, _ := strconv.Atoi(os.Getenv(key))
	return result
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

func GetBoolEnv(key string) bool {
	result, _ := strconv.ParseBool(os.Getenv(key))
	return result
}
