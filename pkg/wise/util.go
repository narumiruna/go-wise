package wise

import "context"

func QueryPrice(ctx context.Context, amount float64, currency, payWith string) (*Price, error) {
	client := NewRestClient()
	req := PriceRequest{
		SourceCurrency: payWith,
		TargetAmount:   amount,
		TargetCurrency: currency,
		ProfileCountry: "GB",
	}
	resp, err := client.QueryPrice(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.VISACreditInBalanceOut()
}

func QueryRate(ctx context.Context, source, target string) (float64, error) {
	client := NewRestClient()
	req := RateRequest{
		Source: source,
		Target: target,
	}
	resp, err := client.QueryRate(ctx, req)
	if err != nil {
		return 0, err
	}
	return resp.Value, nil
}
