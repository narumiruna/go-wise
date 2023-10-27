package main

import (
	"context"
	"fmt"

	"github.com/narumiruna/go-wise/pkg/wise"
)

func main() {
	ctx := context.Background()
	client := wise.NewRestClient()

	sources := []string{"GBP", "NOK", "EUR"}
	for _, source := range sources {
		price, err := client.QueryPrice(ctx, source, 1000, "USD")
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
