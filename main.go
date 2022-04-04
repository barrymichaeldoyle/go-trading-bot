package main

import (
	"go-trading-bot/services"
	"go-trading-bot/tasks"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	tasks.DollarCostAverage()
	services.ConnectToAccountWS()
}
