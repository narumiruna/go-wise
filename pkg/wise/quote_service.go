package wise

import "context"

type QuoteService struct {
	client *Client
}

func (c *Client) NewQuoteService() *QuoteService {
	return &QuoteService{
		client: c,
	}
}

func (s *QuoteService) QueryQuote(ctx context.Context, sourceCurrency string, targetCurrency string, sourceAmount float64, targetAmount float64) (*Quote, error) {
	req := s.client.NewQuoteRequest()
	if sourceCurrency != "" {
		req.SourceCurrency(sourceCurrency)
	}
	if targetCurrency != "" {
		req.TargetCurrency(targetCurrency)
	}
	if sourceAmount != 0 {
		req.SourceAmount(sourceAmount)
	}
	if targetAmount != 0 {
		req.TargetAmount(targetAmount)
	}
	return req.Do(ctx)
}
