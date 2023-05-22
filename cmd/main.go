package main

import (
	"context"
	"fmt"

	"github.com/narumiruna/go-wise/pkg/wise"
)

var sourceCurrencies = []string{"GBP", "EUR"}

func main() {
	for _, sourceCurrency := range sourceCurrencies {
		client := wise.NewRestClient()
		req := wise.PriceRequest{
			TargetAmount:   1000,
			SourceCurrency: sourceCurrency,
			TargetCurrency: "USD",
		}
		resp, err := client.QueryPrice(context.Background(), req)
		if err != nil {
			panic(err)
		}

		data, err := resp.BankTransferInBalanceOut()
		if err != nil {
			panic(err)
		}

		fmt.Printf("%+v\n", data)
	}
}
