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

	req.Header.Add("Content-Type", "application/json")
	return req, nil
}

func (c *Client) NewPriceService() *PriceService {
	return &PriceService{
		client: c,
	}
}

func (c *Client) NewRatesService() *RateService {
	return &RateService{
		client: c,
	}
}
