package cost

import (
	"fmt"

	"github.com/narumiruna/go-wise/pkg/wise"
)

func findPrice(prices []wise.Price, payInMethod wise.PayInMethod, payOutMethod wise.PayOutMethod) (*wise.Price, error) {
	for _, p := range prices {
		if p.PayInMethod == payInMethod && p.PayOutMethod == payOutMethod {
			return &p, nil
		}
	}

	return nil, fmt.Errorf("method not found")
}
