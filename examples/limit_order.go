package examples

import (
	"fmt"
	"go-trading-bot/apis"
)

func LimitOrder() {
	limitOrder, err := apis.PostLimitOrder(apis.PostLimitOrderOptions{
		Side:     "SELL",
		Quantity: "0.001",
		Price:    "1000000",
		Pair:     "BTCZAR",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(limitOrder.Id)
}
