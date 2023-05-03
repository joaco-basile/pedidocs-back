package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}
	uri := os.Getenv("MOGNOURI")
	return uri
}
