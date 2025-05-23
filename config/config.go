package config

import (
	"log"

	"github.com/joho/godotenv"
)


func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConfigureApp (){
	loadEnv()
}