package max

import (
	"context"
	"strconv"

	"github.com/narumiruna/go-wise/pkg/max/maxapi"
)

type BidAsk struct {
	Bid float64
	Ask float64
}

func QueryUSDTTWD() (*BidAsk, error) {
	c := maxapi.NewRestClient()
	t, err := c.QueryTicker(context.Background(), "USDTTWD")
	if err != nil {
		return nil, err
	}

	bid, err := strconv.ParseFloat(t.Buy, 64)
	if err != nil {
		return nil, err
	}

	ask, err := strconv.ParseFloat(t.Sell, 64)
	if err != nil {
		return nil, err
	}

	return &BidAsk{
		Bid: bid,
		Ask: ask,
	}, nil
}
