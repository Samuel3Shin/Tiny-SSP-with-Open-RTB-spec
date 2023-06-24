package common

import (
	"os"
)

type Config struct {
	DSP1_URL      string
	DSP2_URL      string
	SSP_URL       string
	LOGSERVER_URL string
	MONGODB_URL   string
}

func GetConfig() *Config {
	return &Config{
		DSP1_URL:      getEnv("DSP1_URL", "http://127.0.0.1:8081"), // Default to "127.0.0.1" if not set
		DSP2_URL:      getEnv("DSP2_URL", "http://127.0.0.1:8082"),
		SSP_URL:       getEnv("SSP_URL", "http://127.0.0.1:8080"),
		LOGSERVER_URL: getEnv("LOGSERVER_URL", "http://127.0.0.1:8083"),
		MONGODB_URL:   getEnv("MONGODB_URL", "mongodb://127.0.0.1:27017"),
	}
}

// Helper function to get environment variables with a fallback
func getEnv(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}
