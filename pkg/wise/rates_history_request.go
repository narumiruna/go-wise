package wise

import (
	"github.com/c9s/requestgen"
)

//go:generate requestgen -method GET -url "/rates/history" -type RatesHistoryRequest -responseType []Rate
type RatesHistoryRequest struct {
	client requestgen.AuthenticatedAPIClient

	source     string     `param:"source,query"`
	target     string     `param:"target,query"`
	length     int        `param:"length,query"`
	resolution Resolution `param:"resolution,query"`
	unit       Unit       `param:"unit,query"`
}

func (c *RestClient) NewRatesHistoryRequest() *RatesHistoryRequest {
	return &RatesHistoryRequest{
		client: c,
	}
}
