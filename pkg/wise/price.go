package wise

type Price struct {
	PriceSetID           int64                `json:"priceSetId"`
	SourceAmount         float64              `json:"sourceAmount"`
	TargetAmount         float64              `json:"targetAmount"`
	PayInMethod          PayInMethod          `json:"payInMethod"`
	PayOutMethod         PayOutMethod         `json:"payOutMethod"`
	SourceCurrency       string               `json:"sourceCcy"`
	TargetCurrency       string               `json:"targetCcy"`
	Total                float64              `json:"total"`
	VariableFee          float64              `json:"variableFee"`
	VariableFeePercent   float64              `json:"variableFeePercent"`
	SwiftPayoutFlatFee   float64              `json:"swiftPayoutFlatFee"`
	FlatFee              float64              `json:"flatFee"`
	MidRate              float64              `json:"midRate"`
	EcbRate              float64              `json:"ecbRate"`
	EcbRateTimestamp     int64                `json:"ecbRateTimestamp"`
	EcbMarkupPercent     float64              `json:"ecbMarkupPercent"`
	AdditionalFeeDetails AdditionalFeeDetails `json:"additionalFeeDetails"`
}

type AdditionalFeeDetails struct {
	ReceiveInr       interface{} `json:"receiveInr"`
	TaxDetailsByType interface{} `json:"taxDetailsByType"`
	DynamicFxFee     interface{} `json:"dynamicFxFee"`
}

type PriceSlice []Price

func (s PriceSlice) Find(payInMethod PayInMethod, payOutMethod PayOutMethod) (Price, bool) {
	for _, p := range s {
		if p.PayInMethod == payInMethod && p.PayOutMethod == payOutMethod {
			return p, true
		}
	}
	return Price{}, false
}
