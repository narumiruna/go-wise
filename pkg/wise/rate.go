package wise

import (
	"context"

	"github.com/c9s/requestgen"
)

//go:generate requestgen -method GET -url "/rates/live" -type RateRequest -responseType Rate
type RateLiveRequest struct {
	client requestgen.AuthenticatedAPIClient

	source string `param:"source,query"`
	target string `param:"target,query"`
}

func (c *RestClient) NewRateRequest() *RateLiveRequest {
	return &RateLiveRequest{
		client: c,
	}
}

func (c *RestClient) QueryRateLive(ctx context.Context, source, target string) (*Rate, error) {
	req := c.NewRateRequest().Source(source).Target(target)
	return req.Do(ctx)
}

//go:generate requestgen -method GET -url "/rates/history" -type RateHistoryRequest -responseType []Rate
type RateHistoryRequest struct {
	client requestgen.AuthenticatedAPIClient

	source     string     `param:"source,query"`
	target     string     `param:"target,query"`
	length     int        `param:"length,query"`
	resolution Resolution `param:"resolution,query"`
	unit       Unit       `param:"unit,query"`
}

func (c *RestClient) NewRateHistoryRequest() *RateHistoryRequest {
	return &RateHistoryRequest{
		client: c,
	}
}

func (c *RestClient) QueryRateHistory(ctx context.Context, source, target string, length int, resolution Resolution, unit Unit) ([]Rate, error) {
	req := c.NewRateHistoryRequest().Source(source).Target(target).Length(length).Resolution(resolution).Unit(unit)
	return req.Do(ctx)
}

type Rate struct {
	Source string    `json:"source"`
	Target string    `json:"target"`
	Value  float64   `json:"value"`
	Time   Timestamp `json:"time"`
}
