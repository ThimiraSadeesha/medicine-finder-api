package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	MODE    string
	APP_ENV string
	PORT    string
	HOST    string
	DB_TYPE string
	DB_HOST string
	DB_PORT string
	DB_USER string
	DB_PASS string
	DB_NAME string
)

func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, falling back to environment variables")
	}
	MODE = getEnv("MODE", "release")
	APP_ENV = getEnv("APP_ENV", "production")
	PORT = getEnv("PORT", "8080")
	HOST = getEnv("HOST", "127.0.0.1")
	DB_TYPE = getEnv("DB_TYPE", "mysql")
	DB_HOST = getEnv("DB_HOST", "localhost")
	DB_PORT = getEnv("DB_PORT", "3306")
	DB_USER = getEnv("DB_USER", "root")
	DB_PASS = getEnv("DB_PASS", "")
	DB_NAME = getEnv("DB_NAME", "testdb")
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
