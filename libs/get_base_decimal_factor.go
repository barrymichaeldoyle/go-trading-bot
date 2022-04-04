package libs

func GetBaseDecimalFactor(currencyPair string) float64 {
	switch currencyPair {
	case "BTCZAR":
		return 1.0
	case "ETHZAR":
		return 1.0
	case "XRPZAR":
		return 100.0
	case "SOLZAR":
		return 1.0
	case "BTCUSDC":
		return 1.0
	case "USDCZAR":
		return 100.0
	case "BNBZAR":
		return 1.0
	default:
		return 1
	}
}
