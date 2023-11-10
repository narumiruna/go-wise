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
