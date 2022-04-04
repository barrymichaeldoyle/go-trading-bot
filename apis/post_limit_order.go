// https://docs.valr.com/#5beb7328-24ca-4d8a-84f2-6029725ad923
package apis

import (
	"encoding/json"
	"go-trading-bot/config"
	"go-trading-bot/libs"
	"io/ioutil"
	"net/http"
	"strings"
)

type PostLimitOrderOptions struct {
	Side            string `json:"side"`
	Quantity        string `json:"quantity"`
	Price           string `json:"price"`
	Pair            string `json:"pair"`
	PostOnly        bool   `json:"postOnly,omitempty" bson:",omitempty"`
	CustomerOrderId string `json:"customerOrderId,omitempty" bson:",omitempty"`
	TimeInForce     string `json:"timeInForce,omitempty" bson:",omitempty"`
}

type PostLimitOrderResponse struct {
	Id string `json:"id"`
}

func PostLimitOrder(options PostLimitOrderOptions) (*PostLimitOrderResponse, error) {
	baseUrl := "https://api.valr.com"
	endpoint := "/v1/orders/limit"
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
	req.Header.Add("X-VALR-API-KEY", config.API_KEY)
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

	output := &PostLimitOrderResponse{}
	if err := json.Unmarshal(body, output); err != nil {
		return nil, err
	}
	return output, nil
}
