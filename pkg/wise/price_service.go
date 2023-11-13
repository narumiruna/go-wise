package wise

import "context"

type PriceService struct {
	client *Client
}

func (s *PriceService) QueryPrice(ctx context.Context, source string, amount float64, target string) ([]Price, error) {
	req := s.client.NewPriceRequest().SourceCurrency(source).TargetAmount(amount).TargetCurrency(target)
	return req.Do(ctx)
}
