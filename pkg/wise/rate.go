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
	resp, err := req.Do(ctx)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type Rate struct {
	Source string  `json:"source"`
	Target string  `json:"target"`
	Value  float64 `json:"value"`
	Time   int64   `json:"time"`
}

type Resolution string

const (
	HourlyResolution Resolution = "hourly"
	DailyResolution  Resolution = "daily"
)

type Unit string

const (
	DayUnit   Unit = "day"
	MonthUnit Unit = "month"
	YearUnit  Unit = "year"
)
