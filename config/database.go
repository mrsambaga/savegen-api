package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type dbConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

var (
	DBConfig = dbConfig{
		Host:     getENV("DB_HOST", ""),
		User:     getENV("DB_USER", ""),
		Password: getENV("DB_PASSWORD", ""),
		DBName:   getENV("DB_NAME", ""),
		Port:     getENV("DB_PORT", ""),
	}
)

func getENV(key, defaultVal string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	env := os.Getenv(key)
	if env == "" {
		return defaultVal
	}
	return env
}