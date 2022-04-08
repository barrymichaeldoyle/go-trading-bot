package libs

import "strings"

func GetBaseCurrencyFromPair(currencyPair string) string {
	if strings.Contains(currencyPair, "ZAR") {
		return strings.Replace(currencyPair, "ZAR", "", -1)
	}
	return strings.Replace(currencyPair, "USDC", "", -1)
}
