package wise

import (
	"github.com/c9s/requestgen"
)

//go:generate requestgen -method POST -url "/gateway/v3/quotes" -type QuoteRequest -responseType Quote
type QuoteRequest struct {
	client requestgen.AuthenticatedAPIClient

	sourceCurrency string   `param:"sourceCurrency"`
	targetCurrency string   `param:"targetCurrency"`
	sourceAmount   *float64 `param:"sourceAmount"`
	targetAmount   *float64 `param:"targetAmount"`
}

func (c *Client) NewQuoteRequest() *QuoteRequest {
	return &QuoteRequest{
		client: c,
	}
}
