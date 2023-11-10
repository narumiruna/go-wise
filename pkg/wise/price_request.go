package wise

import (
	"github.com/c9s/requestgen"
)

//go:generate requestgen -method GET -url "/gateway/v1/price" -type PriceRequest -responseType []Price
type PriceRequest struct {
	client requestgen.AuthenticatedAPIClient

	sourceAmount   *float64 `param:"sourceAmount"`
	sourceCurrency string   `param:"sourceCurrency"`
	targetAmount   *float64 `param:"targetAmount"`
	targetCurrency string   `param:"targetCurrency"`
	profileID      *int     `param:"profileId"`
	profileCountry *string  `param:"profileCountry"`
	profileType    *string  `param:"profileType"`
	markers        *string  `param:"markers"`
}

func (c *RestClient) NewPriceRequest() *PriceRequest {
	return &PriceRequest{
		client: c,
	}
}
