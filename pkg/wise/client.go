package wise

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/c9s/requestgen"
)

const defaultHTTPTimeout = time.Second * 15
const baseURL = "https://wise.com"

type RestClient struct {
	requestgen.BaseAPIClient
}

func NewRestClient() *RestClient {
	u, err := url.Parse(baseURL)
	if err != nil {
		panic(err)
	}

	return &RestClient{
		BaseAPIClient: requestgen.BaseAPIClient{
			BaseURL: u,
			HttpClient: &http.Client{
				Timeout: defaultHTTPTimeout,
			},
		},
	}
}

func (c *RestClient) NewAuthenticatedRequest(ctx context.Context, method, refURL string, params url.Values, payload interface{}) (*http.Request, error) {
	req, err := c.NewRequest(ctx, method, refURL, params, payload)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (c *RestClient) QueryPrice(ctx context.Context, source string, amount float64, target string) ([]Price, error) {
	req := c.NewPriceRequest().SourceCurrency(source).TargetAmount(amount).TargetCurrency(target)
	return req.Do(ctx)
}

func (c *RestClient) QueryRate(ctx context.Context, source, target string) (*Rate, error) {
	req := c.NewRatesLiveRequest().Source(source).Target(target)
	return req.Do(ctx)
}

func (c *RestClient) QueryRateHistory(ctx context.Context, source, target string, length int, resolution Resolution, unit Unit) ([]Rate, error) {
	req := c.NewRatesHistoryRequest().Source(source).Target(target).Length(length).Resolution(resolution).Unit(unit)
	return req.Do(ctx)
}
