package main

import (
	"context"
	"fmt"

	"github.com/narumiruna/go-wise/pkg/cost"
	"github.com/narumiruna/go-wise/pkg/wise"
)

func main() {
	client := wise.NewClient("")
	ctx := context.Background()
	service := client.NewService()

	prices, err := service.QueryPrice(ctx, "GBP", 1000, "USD")
	if err != nil {
		panic(err)
	}
	price, ok := prices.Find(wise.PayInMethodVisaCredit, wise.PayOutMethodBalance)
	if !ok {
		panic("price not found")
	}

	cost := cost.NewCost(price)
	fmt.Printf("%+v\n", cost)

	rate, err := service.QueryRate(ctx, "GBP", "USD")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", rate)

	rates, err := service.QueryRateHistory(ctx, "GBP", "USD", 14, wise.ResolutionHourly, wise.UnitDay)
	if err != nil {
		panic(err)
	}
	fmt.Printf("len(rates) = %+v\n", len(rates))

	currencies, err := service.QueryCurrency(ctx)
	if err != nil {
		panic(err)
	}
	for _, currency := range currencies {
		fmt.Printf("code: %s, symbol: %s, name: %s\n", currency.Code, currency.Symbol, currency.Name)
	}

	quotes, err := service.QueryQuote(ctx, "GBP", "USD", 0, 1000)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", quotes)
}
