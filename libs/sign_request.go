package libs

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"strings"
)

func SignRequest(timestampString string, verb string, path string, body string) string {
	mac := hmac.New(sha512.New, []byte(GoDotEnvVariable("SECRET")))
	mac.Write([]byte(timestampString))
	mac.Write([]byte(strings.ToUpper(verb)))
	mac.Write([]byte(path))
	mac.Write([]byte(body))
	return hex.EncodeToString(mac.Sum(nil))
}
