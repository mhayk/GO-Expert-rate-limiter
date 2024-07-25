package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	MaxRequestsPerSecond      int
	BlockDurationSeconds      int
	IPMaxRequestsPerSecond    int
	IPBlockDurationSeconds    int
	TokenMaxRequestsPerSecond int
	TokenBlockDurationSeconds int
	RedisHost                 string
	RedisPort                 string
	RedisPassword             string
}

func LoadConfig(envFile string) (*Config, error) {
	if err := godotenv.Load(envFile); err != nil {
		log.Printf("Warning: %v\n", err)
	}

	maxRequestsPerSecond, _ := strconv.Atoi(getEnv("MAX_REQUESTS_PER_SECOND", "100"))
	blockDurationSeconds, _ := strconv.Atoi(getEnv("BLOCK_DURATION_SECONDS", "60"))
	ipMaxRequestsPerSecond, _ := strconv.Atoi(getEnv("IP_MAX_REQUESTS_PER_SECOND", "5"))
	ipBlockDurationSeconds, _ := strconv.Atoi(getEnv("IP_BLOCK_DURATION_SECONDS", "300"))
	tokenMaxRequestsPerSecond, _ := strconv.Atoi(getEnv("TOKEN_MAX_REQUESTS_PER_SECOND", "10"))
	tokenBlockDurationSeconds, _ := strconv.Atoi(getEnv("TOKEN_BLOCK_DURATION_SECONDS", "300"))
	redisHost := getEnv("REDIS_HOST", "localhost")
	redisPort := getEnv("REDIS_PORT", "6379")
	redisPassword := getEnv("REDIS_PASSWORD", "")

	cfg := &Config{
		MaxRequestsPerSecond:      maxRequestsPerSecond,
		BlockDurationSeconds:      blockDurationSeconds,
		IPMaxRequestsPerSecond:    ipMaxRequestsPerSecond,
		IPBlockDurationSeconds:    ipBlockDurationSeconds,
		TokenMaxRequestsPerSecond: tokenMaxRequestsPerSecond,
		TokenBlockDurationSeconds: tokenBlockDurationSeconds,
		RedisHost:                 redisHost,
		RedisPort:                 redisPort,
		RedisPassword:             redisPassword,
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
