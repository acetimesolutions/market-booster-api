package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func ConnectDB() {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// dsn := os.Getenv("dsn")
}
