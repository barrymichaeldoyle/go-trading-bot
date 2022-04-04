package main

import (
	"fmt"
	"go-trading-bot/apis"
)

func main() {
	fmt.Println("Running BTC DCA CRONJOB")
	fmt.Println("Market Order BUY R140 worth of BTC")
	go apis.PostMarketOrder(apis.PostMarketOrderOptions{
		Side:        "BUY",
		QuoteAmount: "140",
		Pair:        "BTCZAR",
	})
}
