package wise

import (
	"context"

	"github.com/c9s/requestgen"
)

//go:generate requestgen -method GET -url "/rates/live" -type RateRequest -responseType Rate
type RateRequest struct {
	client requestgen.AuthenticatedAPIClient

	source string `param:"source,query"`
	target string `param:"target,query"`
}

func (c *RestClient) NewRateRequest() *RateRequest {
	return &RateRequest{
		client: c,
	}
}

func (c *RestClient) QueryRate(ctx context.Context, source, target string) (float64, error) {
	req := c.NewRateRequest().Source(source).Target(target)
	resp, err := req.Do(ctx)
	if err != nil {
		return 0, err
	}
	return resp.Value, nil
}

type Rate struct {
	Source string  `json:"source"`
	Target string  `json:"target"`
	Value  float64 `json:"value"`
	Time   int64   `json:"time"`
}
