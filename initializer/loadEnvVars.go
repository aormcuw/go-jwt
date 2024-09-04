package initializer

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVars() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cannot load the env file")
	}
}
