package libs

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"go-trading-bot/config"
	"strings"
)

func SignRequest(timestampString string, verb string, path string, body string) string {
	mac := hmac.New(sha512.New, []byte(config.API_SECRET))
	mac.Write([]byte(timestampString))
	mac.Write([]byte(strings.ToUpper(verb)))
	mac.Write([]byte(path))
	mac.Write([]byte(body))
	return hex.EncodeToString(mac.Sum(nil))
}
