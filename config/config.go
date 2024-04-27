package config

import (
	"github.com/joho/godotenv"
	"log"
)

func InitConfig() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatalln("Failed to load .env file:", err)
	}
}
