package main

import (
	"context"
	"fmt"

	"github.com/narumiruna/go-wise/pkg/wise"
)

func main() {
	client := wise.NewRestClient()
	ctx := context.Background()

	cost, err := client.NewCost(ctx, "GBP", 1000, "USD")
	if err != nil {
		panic(err)
	}

	fmt.Println(cost.String())

	// rate history
	rates, err := client.QueryRateHistory(ctx, "GBP", "USD", 10, wise.ResolutionHourly, wise.UnitDay)
	if err != nil {
		panic(err)
	}

	for _, rate := range rates {
		fmt.Printf("%+v\n", rate)
	}
}
