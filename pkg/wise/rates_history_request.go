package wise

import (
	"github.com/c9s/requestgen"
)

//go:generate requestgen -method GET -url "/rates/history" -type RatesHistoryRequest -responseType []Rate
type RatesHistoryRequest struct {
	client requestgen.AuthenticatedAPIClient

	source     string     `param:"source"`
	target     string     `param:"target"`
	length     int        `param:"length"`
	resolution Resolution `param:"resolution"`
	unit       Unit       `param:"unit"`
}

func (c *Client) NewRatesHistoryRequest() *RatesHistoryRequest {
	return &RatesHistoryRequest{
		client: c,
	}
}
