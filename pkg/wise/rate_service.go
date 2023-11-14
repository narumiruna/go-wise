package wise

import "context"

type RateService struct {
	client *Client
}

func (s *RateService) QueryRate(ctx context.Context, source, target string) (*Rate, error) {
	req := s.client.NewRateLiveRequest().Source(source).Target(target)
	return req.Do(ctx)
}

func (s *RateService) QueryRateHistory(ctx context.Context, source, target string, length int, resolution Resolution, unit Unit) ([]Rate, error) {
	req := s.client.NewRateHistoryRequest().Source(source).Target(target).Length(length).Resolution(resolution).Unit(unit)
	return req.Do(ctx)
}
