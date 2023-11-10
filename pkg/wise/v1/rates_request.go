package v1

import (
	"github.com/c9s/requestgen"
)

//go:generate requestgen -method GET -url "/v1/rates" -type RatesRequest -responseType []Rate
type RatesRequest struct {
	client requestgen.AuthenticatedAPIClient

	source string  `param:"source"`
	target string  `param:"target"`
	time   *string `param:"time"`
	from   *string `param:"from"`
	to     *string `param:"to"`
	group  *Group  `param:"group"`
}

func (c *RestClient) NewRatesRequest() *RatesRequest {
	return &RatesRequest{
		client: c,
	}
}
