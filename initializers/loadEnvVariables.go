package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	errLocal := godotenv.Load(".localenv")

	if errLocal != nil {
		errEnv := godotenv.Load(".env")

		if errEnv != nil {
			log.Fatal("Error loading .env file")
		}
	}
}
