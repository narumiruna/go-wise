package main

import (
	"context"
	"fmt"

	"github.com/narumiruna/go-wise/pkg/wise"
)

func main() {
	ctx := context.Background()

	sourceCurrencies := []string{"GBP", "NOK", "EUR"}
	for _, sourceCurrency := range sourceCurrencies {
		price, err := wise.QueryPrice(ctx, sourceCurrency, 1000, "USD")
		if err != nil {
			panic(err)
		}
		cost, err := wise.NewCost(ctx, price)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", cost)
	}
}
