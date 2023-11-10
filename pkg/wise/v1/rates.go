package v1

import (
	"github.com/c9s/requestgen"
)

type Rate struct {
	Rate   float64 `json:"rate"`
	Target string  `json:"target"`
	Source string  `json:"source"`
	Time   Time    `json:"time"`
}

//go:generate requestgen -method GET -url "/v1/rates" -type RateRequest -responseType []Rate
type RateRequest struct {
	client requestgen.AuthenticatedAPIClient

	source string  `param:"source,query"`
	target string  `param:"target,query"`
	time   *string `param:"time,query"`
	from   *string `param:"from,query"`
	to     *string `param:"to,query"`
	group  *Group  `param:"group,query"`
}

func (c *RestClient) NewRateRequest() *RateRequest {
	return &RateRequest{
		client: c,
	}
}
