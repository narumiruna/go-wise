package main

import (
	"context"
	"fmt"

	"github.com/narumiruna/go-wise/pkg/cost"
	"github.com/narumiruna/go-wise/pkg/wise"
)

func main() {
	client := wise.NewClient()
	ctx := context.Background()

	prices, err := client.NewPriceService().QueryPrice(ctx, "GBP", 1000, "USD")
	if err != nil {
		panic(err)
	}
	price, ok := prices.Find(wise.PayInMethodVisaCredit, wise.PayOutMethodBalance)
	if !ok {
		panic("price not found")
	}

	cost := cost.NewCost(price)
	fmt.Printf("%+v\n", cost)

	rate, err := client.NewRatesService().QueryRate(ctx, "GBP", "USD")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", rate)

	rates, err := client.NewRatesService().QueryRateHistory(ctx, "GBP", "USD", 14, wise.ResolutionHourly, wise.UnitDay)
	if err != nil {
		panic(err)
	}
	fmt.Printf("len(rates) = %+v\n", len(rates))
}
