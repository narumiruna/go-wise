package wise

import (
	"context"

	"github.com/c9s/requestgen"
)

//go:generate requestgen -method GET -url "/gateway/v1/price" -type PriceRequest -responseType []Price
type PriceRequest struct {
	client requestgen.AuthenticatedAPIClient

	sourceAmount   *float64 `param:"sourceAmount,query"`
	sourceCurrency string   `param:"sourceCurrency,query"`
	targetAmount   *float64 `param:"targetAmount,query"`
	targetCurrency string   `param:"targetCurrency,query"`
	profileID      *int     `param:"profileId,query"`
	profileCountry *string  `param:"profileCountry,query"`
	profileType    *string  `param:"profileType,query"`
	markers        *string  `param:"markers,query"`
}

func (c *RestClient) NewPriceRequest() *PriceRequest {
	return &PriceRequest{
		client: c,
	}
}

func (c *RestClient) QueryPrice(ctx context.Context, source string, amount float64, target string) ([]Price, error) {
	req := c.NewPriceRequest().SourceCurrency(source).TargetAmount(amount).TargetCurrency(target)
	return req.Do(ctx)
}

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
