// https://docs.valr.com/#af083ac6-0514-4979-9bab-f599ea1bed4f
package apis

import (
	"encoding/json"
	"go-trading-bot/libs"
	"io/ioutil"
	"net/http"
	"time"
)

type GetCurrentApiKeyInfoResponse struct {
	Label        string    `json:"label"`
	Permissions  []string  `json:"permissions"`
	AddedAt      time.Time `json:"addedAt"`
	IsSubAccount bool      `json:"isSubAccount"`
}

func GetCurrentApiKeyInfo() (*GetCurrentApiKeyInfoResponse, error) {
	baseUrl := "https://api.valr.com"
	endpoint := "/v1/account/api-keys/current"
	method := "GET"

	req, err := http.NewRequest(method, baseUrl+endpoint, nil)
	if err != nil {
		return nil, err
	}

	timestamp := libs.GetCurrentTimestampString()
	signature := libs.SignRequest(timestamp, method, endpoint, "")
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

	output := &GetCurrentApiKeyInfoResponse{}
	if err := json.Unmarshal(body, output); err != nil {
		return nil, err
	}
	return output, nil
}
