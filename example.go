package main

import (
	"context"
	"fmt"

	"github.com/narumiruna/go-wise/pkg/wise"
)

func main() {
	ctx := context.Background()
	s := wise.NewClient().NewService()
	cost, err := s.NewCost(ctx, "GBP", 1000, "USD")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", cost)
}
