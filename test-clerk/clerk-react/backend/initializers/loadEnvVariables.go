package initializers

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

// LoadEnvVariables loads environment variables from .env file
func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
	fmt.Println("Environment variables loaded successfully")
}
