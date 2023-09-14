package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load(".env", ".localenv")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
