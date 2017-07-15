package bittrex

// Used in getmarkethistory
type Trade struct {
	ID        int     `json:"Id"`
	OrderUuid string  `json:"OrderUuid"`
	Timestamp jTime   `json:"TimeStamp"`
	Quantity  float64 `json:"Quantity"`
	Price     float64 `json:"Price"`
	Total     float64 `json:"Total"`
	FillType  string  `json:"FillType"`
	OrderType string  `json:"OrderType"`
}
