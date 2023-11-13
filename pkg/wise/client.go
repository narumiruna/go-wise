package wise

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/c9s/requestgen"
)

const (
	defaultHTTPTimeout = time.Second * 15
	defaultBaseURL     = "https://wise.com"
)

type Client struct {
	requestgen.BaseAPIClient
}

func NewClient() *Client {
	u, err := url.Parse(defaultBaseURL)
	if err != nil {
		panic(err)
	}

	return &Client{
		BaseAPIClient: requestgen.BaseAPIClient{
			BaseURL: u,
			HttpClient: &http.Client{
				Timeout: defaultHTTPTimeout,
			},
		},
	}
}

func (c *Client) NewAuthenticatedRequest(ctx context.Context, method, refURL string, params url.Values, payload interface{}) (*http.Request, error) {
	req, err := c.NewRequest(ctx, method, refURL, params, payload)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (c *Client) QueryPrice(ctx context.Context, source string, amount float64, target string) ([]Price, error) {
	req := c.NewPriceRequest().SourceCurrency(source).TargetAmount(amount).TargetCurrency(target)
	return req.Do(ctx)
}

func (c *Client) QueryRate(ctx context.Context, source, target string) (*Rate, error) {
	req := c.NewRatesLiveRequest().Source(source).Target(target)
	return req.Do(ctx)
}

func (c *Client) QueryRateHistory(ctx context.Context, source, target string, length int, resolution Resolution, unit Unit) ([]Rate, error) {
	req := c.NewRatesHistoryRequest().Source(source).Target(target).Length(length).Resolution(resolution).Unit(unit)
	return req.Do(ctx)
}
