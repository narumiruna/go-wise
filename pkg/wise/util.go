package wise

import "fmt"

func findPrice(prices []Price, payInMethod, payOutMethod string) (*Price, error) {
	for _, p := range prices {
		if p.PayInMethod == payInMethod && p.PayOutMethod == payOutMethod {
			return &p, nil
		}
	}

	return nil, fmt.Errorf("method not found")
}
