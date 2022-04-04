package examples

import (
	"fmt"
	"go-trading-bot/apis"
)

func ApiKeyInfo() {
	apiKeyInfo, err := apis.GetCurrentApiKeyInfo()
	if err != nil {
		panic(err)
	}
	fmt.Println(apiKeyInfo.Label)
}
