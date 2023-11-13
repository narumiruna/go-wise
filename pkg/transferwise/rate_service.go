package transferwise

import (
	"context"
	"time"
)

type RateService struct {
	client *Client
}

func (s *RateService) QueryRate(ctx context.Context, source string, target string) ([]Rate, error) {
	req := s.client.NewRatesRequest().Source(source).Target(target)
	return req.Do(ctx)
}

func (s *RateService) QueryRateHistory(ctx context.Context, source string, target string, from time.Time, to time.Time, gorup Group) ([]Rate, error) {
	req := s.client.NewRatesRequest().Source(source).Target(target).From(Time(from)).To(Time(to)).Group(gorup)
	return req.Do(ctx)
}
