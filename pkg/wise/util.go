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

func QueryMidRate(ctx context.Context, from, to string) (float64, error) {
	client := NewRestClient()
	req := PriceRequest{
		SourceCurrency: from,
		TargetAmount:   2000,
		TargetCurrency: to,
		ProfileCountry: "TW",
	}
	resp, err := client.QueryPrice(ctx, req)
	if err != nil {
		return 0, err
	}

	price, err := resp.VISACreditInBalanceOut()
	if err != nil {
		return 0, err
	}

	return price.MidRate, nil
}
