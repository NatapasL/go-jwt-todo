package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/NatapasL/go-jwt-todo/infra"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	infra.StartServer()
}
