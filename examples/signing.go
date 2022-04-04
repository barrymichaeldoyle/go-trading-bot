package examples

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"strings"
)

func getSignRequest() string {
	timestampString := "1558014486185"
	verb := "GET"
	path := "/v1/account/balances"
	secret := "4961b74efac86b25cce8fbe4c9811c4c7a787b7a5996660afcc2e287ad864363"
	mac := hmac.New(sha512.New, []byte(secret))
	mac.Write([]byte(timestampString))
	mac.Write([]byte(strings.ToUpper(verb)))
	mac.Write([]byte(path))
	return hex.EncodeToString(mac.Sum(nil))
}

func postSignRequest() string {
	timestampString := "1558017528946"
	verb := "POST"
	path := "/v1/orders/market"
	secret := "4961b74efac86b25cce8fbe4c9811c4c7a787b7a5996660afcc2e287ad864363"
	body := `{"customerOrderId":"ORDER-000001","pair":"BTCZAR","side":"BUY","quoteAmount":"80000"}`
	mac := hmac.New(sha512.New, []byte(secret))
	mac.Write([]byte(timestampString))
	mac.Write([]byte(strings.ToUpper(verb)))
	mac.Write([]byte(path))
	mac.Write([]byte(body))
	return hex.EncodeToString(mac.Sum(nil))
}

func TestMethods() {
	getSignature := getSignRequest()
	if getSignature == "9d52c181ed69460b49307b7891f04658e938b21181173844b5018b2fe783a6d4c62b8e67a03de4d099e7437ebfabe12c56233b73c6a0cc0f7ae87e05f6289928" {
		fmt.Println("Get Success")
	} else {
		fmt.Println("YOU SUCK at Get!")
	}

	postSignature := postSignRequest()
	if postSignature == "be97d4cd9077a9eea7c4e199ddcfd87408cb638f2ec2f7f74dd44aef70a49fdc49960fd5de9b8b2845dc4a38b4fc7e56ef08f042a3c78a3af9aed23ca80822e8" {
		fmt.Println("Post Success")
	} else {
		fmt.Println("YOU SUCK at Post!")
	}
}
