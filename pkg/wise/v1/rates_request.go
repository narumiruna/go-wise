package v1

import (
	"github.com/c9s/requestgen"
)

//go:generate requestgen -method GET -url "/v1/rates" -type RatesRequest -responseType []Rate
type RatesRequest struct {
	client requestgen.AuthenticatedAPIClient

	source string  `param:"source,query"`
	target string  `param:"target,query"`
	time   *string `param:"time,query"`
	from   *string `param:"from,query"`
	to     *string `param:"to,query"`
	group  *Group  `param:"group,query"`
}

func (c *RestClient) NewRatesRequest() *RatesRequest {
	return &RatesRequest{
		client: c,
	}
}
