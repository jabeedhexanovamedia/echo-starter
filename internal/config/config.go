package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Config is the global typed configuration struct
type Config struct {
	Server      ServerConfig
	Database    DatabaseConfig
	Logging     LoggingConfig
	RateLimiter RateLimiterConfig
}

// ServerConfig holds server-related configs
type ServerConfig struct {
	Port string
	Env  string
}

// DatabaseConfig holds DB connection info
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
}

// LoggingConfig holds log settings
type LoggingConfig struct {
	Level string
}

// RateLimiterConfig holds rate limiting settings
type RateLimiterConfig struct {
	Requests int
	Duration time.Duration
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	// Parse DB port
	dbPort, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
	if err != nil {
		log.Fatalf("Invalid DB_PORT: %v", err)
	}

	// Parse Rate Limiter
	reqs, err := strconv.Atoi(getEnv("RATE_LIMIT_REQUESTS", "20"))
	if err != nil {
		log.Fatalf("Invalid RATE_LIMIT_REQUESTS: %v", err)
	}

	duration, err := time.ParseDuration(getEnv("RATE_LIMIT_DURATION", "1s"))
	if err != nil {
		log.Fatalf("Invalid RATE_LIMIT_DURATION: %v", err)
	}

	return &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "8080"),
			Env:  getEnv("ENV", "development"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     dbPort,
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "password"),
			Name:     getEnv("DB_NAME", "echo_starter"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		Logging: LoggingConfig{
			Level: getEnv("LOG_LEVEL", "debug"),
		},
		RateLimiter: RateLimiterConfig{
			Requests: reqs,
			Duration: duration,
		},
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		return val
	}

	return fallback
}
