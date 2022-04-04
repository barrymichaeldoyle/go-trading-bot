package main

import (
	"go-trading-bot/services"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	services.ConnectToAccountWS()
}
