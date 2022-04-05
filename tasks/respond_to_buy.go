package tasks

import (
	"fmt"
	"go-trading-bot/apis"
	"go-trading-bot/libs"
	"go-trading-bot/types"
	"math"
	"strconv"
)

func RespondToBuy(data types.OrderStatusUpdateData) string {
	executedPriceFloat, err := strconv.ParseFloat(data.ExecutedPrice, 64)
	if err != nil {
		panic(err)
	}
	sellOffPrice := executedPriceFloat * 1.005
	decimalFactor := libs.GetBaseDecimalFactor(data.CurrencyPair)
	newPrice := strconv.FormatFloat(
		math.Ceil(sellOffPrice*decimalFactor)/decimalFactor, 'f', -1, 64)
	out, err := apis.PostLimitOrder(apis.PostLimitOrderOptions{
		Side:     "SELL",
		Quantity: data.ExecutedQuantity,
		Price:    newPrice,
		Pair:     data.CurrencyPair,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Created SELL order for " +
		data.ExecutedQuantity +
		libs.GetBaseCurrencyFromPair(data.CurrencyPair) +
		" @" +
		libs.GetFiatSymbolFromCurrencyPair(data.CurrencyPair) +
		newPrice,
	)

	return out.Id
}
