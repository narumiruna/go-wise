package wise

import (
	"github.com/c9s/requestgen"
)

//go:generate requestgen -method GET -url "/rates/live" -type RateLiveRequest -responseType Rate
type RateLiveRequest struct {
	client requestgen.AuthenticatedAPIClient

	source string `param:"source"`
	target string `param:"target"`
}

func (c *Client) NewRateLiveRequest() *RateLiveRequest {
	return &RateLiveRequest{
		client: c,
	}
}
