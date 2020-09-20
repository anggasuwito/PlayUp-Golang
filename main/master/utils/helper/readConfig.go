package helper

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// readEnv read config form .env file
func ReadEnv(key string) string {
	err := godotenv.Load("config.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

