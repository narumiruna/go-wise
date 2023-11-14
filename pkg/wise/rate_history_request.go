package wise

import (
	"github.com/c9s/requestgen"
)

//go:generate requestgen -method GET -url "/rates/history" -type RateHistoryRequest -responseType []Rate
type RateHistoryRequest struct {
	client requestgen.AuthenticatedAPIClient

	source     string     `param:"source"`
	target     string     `param:"target"`
	length     int        `param:"length"`
	resolution Resolution `param:"resolution"`
	unit       Unit       `param:"unit"`
}

func (c *Client) NewRateHistoryRequest() *RateHistoryRequest {
	return &RateHistoryRequest{
		client: c,
	}
}
