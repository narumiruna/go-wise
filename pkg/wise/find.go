package wise

import "fmt"

func FindPriceByMethod(prices []Price, payInMethod PayInMethod, payOutMethod PayOutMethod) (Price, error) {
	for _, p := range prices {
		if p.PayInMethod == payInMethod && p.PayOutMethod == payOutMethod {
			return p, nil
		}
	}
	return Price{}, fmt.Errorf("price not found")
}
