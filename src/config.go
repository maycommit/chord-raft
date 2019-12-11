package sdproject

import (
	"os"
	"strconv"
)

const projectPath = "~/.sdproject"
const logsPath = "./logs"
const snapsPath = "./snaps"

type Config struct {
	ChordSize              int
	MemorySize             int
	GrpcTimeout            int
	StabilizeInterval      int
	FixFingerTableInterval int
	FlushMemoryInterval    int
	SnapshotTrigger        int
	LogPath                string
	SnapPath               string
}

func NewConfig() *Config {
	newConfig := &Config{}
	newConfig.ChordSize, _ = strconv.Atoi(newConfig.getEnv("CHORD_SIZE", "3"))
	newConfig.MemorySize, _ = strconv.Atoi(newConfig.getEnv("MEMORY_SIZE", "2"))
	newConfig.GrpcTimeout, _ = strconv.Atoi(newConfig.getEnv("GRPC_TIMEOUT", "8"))
	newConfig.StabilizeInterval, _ = strconv.Atoi(newConfig.getEnv("STABILIZE_INTERVAL", "100"))
	newConfig.FixFingerTableInterval, _ = strconv.Atoi(newConfig.getEnv("FIX_FINGER_TABLE_INTERVAL", "100"))
	newConfig.FlushMemoryInterval, _ = strconv.Atoi(newConfig.getEnv("FLUSH_MEMORY_INTERVAL", "100"))
	newConfig.SnapshotTrigger, _ = strconv.Atoi(newConfig.getEnv("SNAPSHOT_TRIGGER", "3"))
	newConfig.LogPath = logsPath
	newConfig.SnapPath = snapsPath

	return newConfig
}

func (config *Config) getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}
