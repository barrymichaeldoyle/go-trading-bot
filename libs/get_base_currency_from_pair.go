package libs

func GetBaseCurrencyFromPair(currencyPair string) string {
	return currencyPair[:3]
}
