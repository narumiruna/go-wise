package v1

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/c9s/requestgen"
)

const (
	defaultHTTPTimeout = time.Second * 15
	productionURL      = "https://api.transferwise.com"
	sandboxURL         = "https://api.sandbox.transferwise.tech"
)

type RestClient struct {
	requestgen.BaseAPIClient

	token string
}

func NewRestClient() *RestClient {
	u, err := url.Parse(productionURL)
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

func (c *RestClient) Auth(token string) {
	c.token = token
}

func (c *RestClient) NewAuthenticatedRequest(ctx context.Context, method, refURL string, params url.Values, payload interface{}) (*http.Request, error) {
	req, err := c.NewRequest(ctx, method, refURL, params, payload)
	if err != nil {
		return nil, err
	}

	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}
	return req, nil
}

func (c *RestClient) QueryRate(ctx context.Context, source string, target string) ([]Rate, error) {
	req := c.NewRatesRequest().Source(source).Target(target)
	return req.Do(ctx)
}

func (c *RestClient) QueryRateHistory(ctx context.Context, source string, target string, from time.Time, to time.Time, gorup Group) ([]Rate, error) {
	req := c.NewRatesRequest().Source(source).Target(target).From(Time(from)).To(Time(to)).Group(gorup)
	return req.Do(ctx)
}
