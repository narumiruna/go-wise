package wise

import "context"

func QueryPrice(amount float64, currency, payWith string) (*Price, error) {
	client := NewRestClient()
	request := PriceRequest{
		SourceCurrency: payWith,
		TargetAmount:   amount,
		TargetCurrency: currency,
		ProfileCountry: "GB",
	}
	req, err := client.QueryPrice(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return req.BankTransferInBalanceOut()
}
