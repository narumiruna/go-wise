package wise

type PriceRequest struct {
	SourceAmount   float64 `url:"sourceAmount,omitempty"`
	TargetAmount   float64 `url:"targetAmount,omitempty"`
	SourceCurrency string  `url:"sourceCurrency"`
	TargetCurrency string  `url:"targetCurrency"`
	ProfileID      int     `url:"profileId,omitempty"`
	ProfileCountry string  `url:"profileCountry,omitempty"`
	ProfileType    string  `url:"profileType,omitempty"`
	Markers        string  `url:"markers,omitempty"`
}

type Price struct {
	PriceSetID           int              `json:"priceSetId"`
	SourceAmount         float64          `json:"sourceAmount"`
	TargetAmount         float64          `json:"targetAmount"`
	PayInMethod          string           `json:"payInMethod"`
	PayOutMethod         string           `json:"payOutMethod"`
	SourceCurrency       string           `json:"sourceCcy"`
	TargetCurrency       string           `json:"targetCcy"`
	Total                float64          `json:"total"`
	VariableFee          float64          `json:"variableFee"`
	VariableFeePercent   float64          `json:"variableFeePercent"`
	SwiftPayoutFlatFee   float64          `json:"swiftPayoutFlatFee"`
	FlatFee              float64          `json:"flatFee"`
	MidRate              float64          `json:"midRate"`
	EcbRate              float64          `json:"ecbRate"`
	EcbRateTimestamp     int              `json:"ecbRateTimestamp"`
	EcbMarkupPercent     interface{}      `json:"ecbMarkupPercent"`
	AdditionalFeeDetails TaxDetailsByType `json:"additionalFeeDetails"`
}

type TaxDetailsByType struct {
	ReceiveInr       interface{} `json:"receiveInr"`
	TaxDetailsByType interface{} `json:"taxDetailsByType"`
}

type RateRequest struct {
	Source string `url:"source"`
	Target string `url:"target"`
}

type Rate struct {
	Source string  `json:"source"`
	Target string  `json:"target"`
	Value  float64 `json:"value"`
	Time   int64   `json:"time"`
}
