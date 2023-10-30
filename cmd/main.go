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

	fmt.Printf("%+v\n", cost)
}
