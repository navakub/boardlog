package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort   string
	DBHost    string
	DBUser    string
	DBPass    string
	DBName    string
	DBPort    string
	DBSSLMode string
}

func LoadConfig() *Config {
	_ = godotenv.Load() // loads from .env (optional)

	cfg := &Config{
		AppPort:   GetEnv("APP_PORT", "8080"),
		DBHost:    GetEnv("DB_HOST", "localhost"),
		DBUser:    GetEnv("DB_USER", "postgres"),
		DBPass:    GetEnv("DB_PASS", ""),
		DBName:    GetEnv("DB_NAME", "mydb"),
		DBPort:    GetEnv("DB_PORT", "5432"),
		DBSSLMode: GetEnv("DB_SSLMODE", "disable"),
	}

	return cfg
}

func GetEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}
