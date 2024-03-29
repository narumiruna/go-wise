package transferwise

import (
	"github.com/c9s/requestgen"
)

//go:generate requestgen -method GET -url "/v1/rates" -type RateRequest -responseType []Rate
type RateRequest struct {
	client requestgen.AuthenticatedAPIClient

	source string `param:"source"`
	target string `param:"target"`
	time   *Time  `param:"time"`
	from   *Time  `param:"from"`
	to     *Time  `param:"to"`
	group  *Group `param:"group"`
}

func (c *Client) NewRateRequest() *RateRequest {
	return &RateRequest{
		client: c,
	}
}
