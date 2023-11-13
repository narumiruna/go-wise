package transferwise

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/c9s/requestgen"
)

const (
	defaultHTTPTimeout = time.Second * 15
	defaultBaseURL     = "https://api.transferwise.com"
)

type Client struct {
	requestgen.BaseAPIClient

	token string
}

func NewRestClient() *Client {
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

func (c *Client) Auth(token string) {
	c.token = token
}

func (c *Client) NewAuthenticatedRequest(ctx context.Context, method, refURL string, params url.Values, payload interface{}) (*http.Request, error) {
	req, err := c.NewRequest(ctx, method, refURL, params, payload)
	if err != nil {
		return nil, err
	}

	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}
	return req, nil
}

func (c *Client) QueryRate(ctx context.Context, source string, target string) ([]Rate, error) {
	req := c.NewRatesRequest().Source(source).Target(target)
	return req.Do(ctx)
}

func (c *Client) QueryRateHistory(ctx context.Context, source string, target string, from time.Time, to time.Time, gorup Group) ([]Rate, error) {
	req := c.NewRatesRequest().Source(source).Target(target).From(Time(from)).To(Time(to)).Group(gorup)
	return req.Do(ctx)
}
