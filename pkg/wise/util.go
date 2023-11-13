package wise

import "fmt"

func FindPrice(prices []Price, payInMethod PayInMethod, payOutMethod PayOutMethod) (*Price, error) {
	for _, p := range prices {
		if p.PayInMethod == payInMethod && p.PayOutMethod == payOutMethod {
			return &p, nil
		}
	}

	return nil, fmt.Errorf("method not found")
}
