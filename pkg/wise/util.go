package wise

import (
	"context"
	"fmt"
	"strings"
)

func findPrice(prices []Price, payInMethod, payOutMethod string) (*Price, error) {
	for _, p := range prices {
		if p.PayInMethod == payInMethod && p.PayOutMethod == payOutMethod {
			return &p, nil
		}
	}

	return nil, fmt.Errorf("method not found")
}

func QueryPrice(ctx context.Context, source string, amount float64, target string) (*Price, error) {
	client := NewRestClient()
	req := PriceRequest{
		SourceCurrency: strings.ToUpper(source),
		TargetAmount:   amount,
		TargetCurrency: strings.ToUpper(target),
		ProfileCountry: "GB",
	}
	resp, err := client.QueryPrice(ctx, req)
	if err != nil {
		return nil, err
	}
	return findPrice(resp, "VISA_CREDIT", "BALANCE")
}

func QueryRate(ctx context.Context, source, target string) (float64, error) {
	client := NewRestClient()
	req := RateRequest{
		Source: strings.ToUpper(source),
		Target: strings.ToUpper(target),
	}
	resp, err := client.QueryRate(ctx, req)
	if err != nil {
		return 0, err
	}
	return resp.Value, nil
}
