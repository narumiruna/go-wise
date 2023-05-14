package main

import (
	"context"
	"fmt"

	"github.com/narumiruna/go-wise/pkg/wise"
)

func main() {
	client := wise.NewRestClient()
	req := wise.PriceRequest{
		TargetAmount:   1000,
		SourceCurrency: "GBP",
		TargetCurrency: "USD",
	}
	resp, err := client.QueryPrice(context.Background(), req)
	if err != nil {
		panic(err)
	}

	data, err := resp.GooglePayInBalanceOut()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", data)
}
