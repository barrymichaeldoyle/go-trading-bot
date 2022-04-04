package libs

import (
	"strconv"
	"time"
)

func GetCurrentTimestampString() string {
	timestamp := time.Now()
	return strconv.FormatInt(timestamp.UnixNano()/1000000, 10)
}
