package wise

import (
	"github.com/c9s/requestgen"
)

//go:generate requestgen -method GET -url "/rates/live" -type RatesLiveRequest -responseType Rate
type RatesLiveRequest struct {
	client requestgen.AuthenticatedAPIClient

	source string `param:"source"`
	target string `param:"target"`
}

func (c *RestClient) NewRatesLiveRequest() *RatesLiveRequest {
	return &RatesLiveRequest{
		client: c,
	}
}
