package internal

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config : get env value
func Config(key string) string {
	return os.Getenv(key)
}

// LoadEnv initially load env
func LoadEnv() {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}
}
