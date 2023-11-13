package wise

import "context"

type RatesService struct {
	client *Client
}

func (s *RatesService) QueryRate(ctx context.Context, source, target string) (*Rate, error) {
	req := s.client.NewRatesLiveRequest().Source(source).Target(target)
	return req.Do(ctx)
}

func (s *RatesService) QueryRateHistory(ctx context.Context, source, target string, length int, resolution Resolution, unit Unit) ([]Rate, error) {
	req := s.client.NewRatesHistoryRequest().Source(source).Target(target).Length(length).Resolution(resolution).Unit(unit)
	return req.Do(ctx)
}
