// https://docs.valr.com/#8d9252e1-ee27-495e-86ed-57458bdafd19 (orderId)
// https://docs.valr.com/#87c78a99-c94c-4b16-a986-5957a62a66fc (cusomerOrderId)
package apis

import (
	"encoding/json"
	"errors"
	"go-trading-bot/config"
	"go-trading-bot/libs"
	"io/ioutil"
	"net/http"
	"time"
)

func buildEndpointFromOptions(options GetOrderStatusOptions) (string, error) {
	if options.CurrencyPair == "" {
		return "", errors.New("currencyPair is required")
	}
	if (options.OrderId == "" && options.CustomerOrderId == "") || (options.OrderId != "" && options.CustomerOrderId != "") {
		return "", errors.New("Either orderId or customerOrderId must be specified")
	}
	endpoint := "/v1/orders/" + options.CurrencyPair
	if options.OrderId != "" {
		endpoint += "/orderid/" + options.OrderId
	}
	if options.CustomerOrderId != "" {
		endpoint += "/customerorderid/" + options.CustomerOrderId
	}
	return endpoint, nil
}

type GetOrderStatusOptions struct {
	CurrencyPair    string `json:"currencyPair"`
	OrderId         string `json:"orderId,omitempty" bson:",omitempty"`
	CustomerOrderId string `json:"customerOrderId,omitempty" bson:",omitempty"`
}

type GetOrderStatusResponse struct {
	OrderId           string    `json:"orderId"`
	OrderStatusType   string    `json:"orderStatusType"`
	CurrencyPair      string    `json:"currencyPair"`
	OriginalPrice     string    `json:"originalPrice"`
	RemainingQuantity string    `json:"remainingQuantity"`
	OriginalQuantity  string    `json:"originalQuantity"`
	OrderSide         string    `json:"orderSide"`
	OrderType         string    `json:"orderType"`
	FailedReason      string    `json:"failedReason,omitempty" bson:",omitempty"`
	CustomerOrderId   string    `json:"customerOrderId"`
	OrderUpdatedAt    time.Time `json:"orderUpdatedAt"`
	OrderCreatedAt    time.Time `json:"orderCreatedAt"`
}

func GetOrderStatus(options GetOrderStatusOptions) (*GetOrderStatusResponse, error) {
	baseUrl := "https://api.valr.com"
	method := "GET"
	endpoint, err := buildEndpointFromOptions(options)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, baseUrl+endpoint, nil)
	if err != nil {
		return nil, err
	}
	println(endpoint)

	timestamp := libs.GetCurrentTimestampString()
	signature := libs.SignRequest(timestamp, method, endpoint, "")
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

	println(body)
	output := &GetOrderStatusResponse{}
	if err := json.Unmarshal(body, output); err != nil {
		return nil, err
	}
	println("RECKED 555")

	return output, nil
}
