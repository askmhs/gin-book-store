package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string
}

var AppConfig *Config

func LoadConfig() {
	err := godotenv.Load()

	if err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}

	AppConfig = &Config{
		AppPort: getEnv("APP_PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
