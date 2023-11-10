package wise

type Price struct {
	PriceSetID           int               `json:"priceSetId"`
	SourceAmount         float64           `json:"sourceAmount"`
	TargetAmount         float64           `json:"targetAmount"`
	PayInMethod          string            `json:"payInMethod"`
	PayOutMethod         string            `json:"payOutMethod"`
	SourceCurrency       string            `json:"sourceCcy"`
	TargetCurrency       string            `json:"targetCcy"`
	Total                float64           `json:"total"`
	VariableFee          float64           `json:"variableFee"`
	VariableFeePercent   float64           `json:"variableFeePercent"`
	SwiftPayoutFlatFee   float64           `json:"swiftPayoutFlatFee"`
	FlatFee              float64           `json:"flatFee"`
	MidRate              float64           `json:"midRate"`
	EcbRate              float64           `json:"ecbRate"`
	EcbRateTimestamp     int               `json:"ecbRateTimestamp"`
	EcbMarkupPercent     interface{}       `json:"ecbMarkupPercent"`
	AdditionalFeeDetails *TaxDetailsByType `json:"additionalFeeDetails"`
}

type TaxDetailsByType struct {
	ReceiveInr       interface{} `json:"receiveInr"`
	TaxDetailsByType interface{} `json:"taxDetailsByType"`
}
