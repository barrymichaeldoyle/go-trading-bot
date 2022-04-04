package examples

import (
	"encoding/json"
	"fmt"
	"go-trading-bot/apis"
)

func OrderStatus() {
	getOrderStatus, err := apis.GetOrderStatus(apis.GetOrderStatusOptions{
		CurrencyPair: "BTCZAR",
		OrderId:      "33db38eb-f289-4808-bf22-56e65d700ece",
	})
	if err != nil {
		panic(err)
	}
	out, err := json.Marshal(getOrderStatus)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}
