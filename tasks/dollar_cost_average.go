package tasks

import (
	"fmt"
	"go-trading-bot/apis"
	"time"
)

func DollarCostAverage() {
	fmt.Println("Will run dollar cost average task in 24 hours to buy R140 worth of BTC - make sure weekly R1000 is deposited")
	for {
		<-time.After(time.Hour * 24)
		go apis.PostMarketOrder(apis.PostMarketOrderOptions{
			Side:        "BUY",
			QuoteAmount: "140",
			Pair:        "BTCZAR",
		})
	}
}
