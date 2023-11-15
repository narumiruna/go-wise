package wise

import (
	"context"
)

type Service struct {
	client *Client
}

func (c *Client) NewService() *Service {
	return &Service{
		client: c,
	}
}

func (s *Service) NewCost(ctx context.Context, source string, amount float64, target string) (*Cost, error) {
	prices, err := s.QueryPrice(ctx, source, amount, target)
	if err != nil {
		return nil, err
	}

	price, err := FindPriceByMethod(prices, PayInMethodVisaCredit, PayOutMethodBalance)
	if err != nil {
		return nil, err
	}

	return NewCost(price), err
}

func (s *Service) QueryPrice(ctx context.Context, source string, amount float64, target string) ([]Price, error) {
	req := s.client.NewPriceRequest().SourceCurrency(source).TargetAmount(amount).TargetCurrency(target)
	return req.Do(ctx)
}

func (s *Service) QueryRate(ctx context.Context, source, target string) (*Rate, error) {
	req := s.client.NewRateLiveRequest().Source(source).Target(target)
	return req.Do(ctx)
}

func (s *Service) QueryRateHistory(ctx context.Context, source, target string, length int, resolution Resolution, unit Unit) ([]Rate, error) {
	req := s.client.NewRateHistoryRequest().Source(source).Target(target).Length(length).Resolution(resolution).Unit(unit)
	return req.Do(ctx)
}

func (s *Service) QueryQuote(ctx context.Context, sourceCurrency string, targetCurrency string, sourceAmount float64, targetAmount float64) (*Quote, error) {
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

func (s *Service) QueryCurrency(ctx context.Context) ([]Currency, error) {
	req := s.client.NewCurrencyRequest()
	return req.Do(ctx)
}
