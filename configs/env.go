package configs 

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	if err := godotenv.Load(); err != nil {
		log.Print("Unable to load .env file in root directory")
	}
	return os.Getenv("MONGO_URI")
}