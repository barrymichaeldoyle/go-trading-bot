package libs

import "strings"

func GetFiatSymbolFromCurrencyPair(currencyPair string) string {
	if strings.Contains(currencyPair, "ZAR") {
		return "R"
	}
	return "$"
}