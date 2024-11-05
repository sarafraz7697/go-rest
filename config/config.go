package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all necessary configuration values for the application
type Config struct {
	Port      string
	SecretKey string

	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

// AppConfig is the globally accessible configuration instance
var AppConfig *Config

// LoadConfig loads environment variables into AppConfig, with validation
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file, relying on system environment variables")
	}

	AppConfig = &Config{
		Port:      getEnvOrFatal("PORT"),
		SecretKey: getEnvOrFatal("SECRET_KEY"),

		DBHost: getEnvOrFatal("DB_HOST"),
		DBUser: getEnvOrFatal("DB_USER"),
		DBPass: getEnvOrFatal("DB_PASS"),
		DBPort: getEnvOrFatal("DB_PORT"),
		DBName: getEnvOrFatal("DB_NAME"),
	}

	return AppConfig
}

// getEnvOrFatal gets an environment variable or exits with a message if it’s missing
func getEnvOrFatal(key string) string {
	value := os.Getenv(key)
	if value == "" {
		fmt.Printf("Environment variable %s is required but not set\n", key)
		os.Exit(1)
	}
	return value
}
