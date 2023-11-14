package wise

import "context"

type CurrencyService struct {
	client *Client
}

func (s *CurrencyService) QueryCurrencies(ctx context.Context, source string, amount float64, target string) ([]Currency, error) {
	req := s.client.NewCurrencyRequest()
	return req.Do(ctx)
}
