package main

import (
	"context"

	"github.com/narumiruna/go-wise/pkg/wise"
	log "github.com/sirupsen/logrus"
)

var sourceCurrencies = []string{"GBP", "NOK", "EUR"}

func main() {
	ctx := context.Background()
	for _, sourceCurrency := range sourceCurrencies {
		price, err := wise.QueryPrice(ctx, 1000, "USD", sourceCurrency)
		if err != nil {
			log.Errorf("query price failed: %v", err)
			continue
		}
		cost, err := wise.NewCost(ctx, price)
		if err != nil {
			log.Errorf("failed to create cost: %v", err)
			continue
		}
		log.Infof("%+v", cost)
	}

}
