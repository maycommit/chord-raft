package sdproject

import (
	"os"
	"strconv"
)

type Config struct {
	ChordSize   int
	MemorySize  int
	GrpcTimeout int
}

func NewConfig() *Config {
	newConfig := &Config{}
	newConfig.ChordSize, _ = strconv.Atoi(newConfig.getEnv("CHORD_SIZE", "3"))
	newConfig.MemorySize, _ = strconv.Atoi(newConfig.getEnv("MEMORY_SIZE", "2"))
	newConfig.GrpcTimeout, _ = strconv.Atoi(newConfig.getEnv("GRPC_TIMEOUT", "8"))

	return newConfig
}

func (config *Config) getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}
