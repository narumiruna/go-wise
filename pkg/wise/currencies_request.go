package wise

import (
	"github.com/c9s/requestgen"
)

//go:generate requestgen -method GET -url "/gateway/v1/currencies" -type CurrencyRequest -responseType []Currency
type CurrencyRequest struct {
	client requestgen.AuthenticatedAPIClient
}

func (c *Client) NewCurrencyRequest() *CurrencyRequest {
	return &CurrencyRequest{
		client: c,
	}
}
