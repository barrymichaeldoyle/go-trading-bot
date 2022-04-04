package types

type OrderStatusUpdateData struct {
	OrderId           string `json:"orderId"`
	OrderStatusType   string `json:"orderStatusType"`
	CurrencyPair      string `json:"currencyPair"`
	OriginalPrice     string `json:"originalPrice"`
	RemainingQuantity string `json:"remainingQuantity"`
	OriginalQuantity  string `json:"originalQuantity"`
	OrderSide         string `json:"orderSide"`
	OrderType         string `json:"orderType"`
	FailedReason      string `json:"failedReason"`
	OrderUpdatedAt    string `json:"orderUpdatedAt"`
	OrderCreatedAt    string `json:"orderCreatedAt"`
	ExecutedPrice     string `json:"executedPrice"`
	ExecutedQuantity  string `json:"executedQuantity"`
	ExecutedFee       string `json:"executedFee"`
}

type OrderStatusUpdate struct {
	Type string `json:"type"`
	Data OrderStatusUpdateData
}
