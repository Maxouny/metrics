package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	PID int32
}

func LoadConfig() Config {
	pid, err := strconv.Atoi(getEnv("PROCESS_PID", "12345"))
	if err != nil {
		log.Fatalf("Ошибка преобразования PID: %v", err)
	}

	return Config{
		PID: int32(pid),
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
