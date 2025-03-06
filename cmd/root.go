package cmd

import (
    "log"

    "github.com/joho/godotenv"

)

func Execute() {
	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
	go startGRPC()
	startGateway()
}