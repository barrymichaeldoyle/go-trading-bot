// https://docs.valr.com/#e1892b20-2b2a-44cf-a67b-1d86def85ec4
package apis

import (
	"encoding/json"
	"go-trading-bot/libs"
	"io/ioutil"
	"net/http"
	"strings"
)

type PostMarketOrderOptions struct {
	Side        string `json:"side"`
	BaseAmount  string `json:"baseAmount,omitempty" bson:",omitempty"`
	QuoteAmount string `json:"quoteAmount,omitempty" bson:",omitempty"`
	Pair        string `json:"pair"`
}

type PostMarketOrderResponse struct {
	Id string `json:"id"`
}

func PostMarketOrder(options PostMarketOrderOptions) (*PostMarketOrderResponse, error) {
	baseUrl := "https://api.valr.com"
	endpoint := "/v1/orders/market"
	method := "POST"

	out, err := json.Marshal(options)
	if err != nil {
		return nil, err
	}

	optionsString := string(out)
	payload := strings.NewReader(optionsString)

	timestamp := libs.GetCurrentTimestampString()
	signature := libs.SignRequest(timestamp, method, endpoint, optionsString)
	req, err := http.NewRequest(method, baseUrl+endpoint, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-VALR-API-KEY", libs.GoDotEnvVariable("API_KEY"))
	req.Header.Add("X-VALR-SIGNATURE", signature)
	req.Header.Add("X-VALR-TIMESTAMP", timestamp)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	output := &PostMarketOrderResponse{}
	if err := json.Unmarshal(body, output); err != nil {
		return nil, err
	}
	return output, nil
}
