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
	sandboxBaseURL     = "https://api.sandbox.transferwise.tech"
)

type Client struct {
	requestgen.BaseAPIClient

	token string
}

func newClient(rawURL, token string) *Client {
	u, err := url.Parse(rawURL)
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
		token: token,
	}
}

func NewClient(token string) *Client {
	return newClient(defaultBaseURL, token)
}

func NewSandBoxClient(token string) *Client {
	return newClient(sandboxBaseURL, token)
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

func (c *Client) NewRateService() *RateService {
	return &RateService{
		client: c,
	}
}
