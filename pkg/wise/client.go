package wise

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

const defaultTimeout = 15 * time.Second
const baseApiUrl = "https://wise.com"

type RestClient struct {
	client  *http.Client
	baseURL *url.URL
}

func NewRestClient() *RestClient {
	u, err := url.Parse(baseApiUrl)
	if err != nil {
		panic(err)
	}

	return &RestClient{
		client: &http.Client{
			Timeout: defaultTimeout,
		},
		baseURL: u,
	}
}

func (c *RestClient) NewRequest(ctx context.Context, method string, refURL string, params url.Values) (*http.Request, error) {
	rel, err := url.Parse(refURL)
	if err != nil {
		return nil, err
	}

	if params != nil {
		rel.RawQuery = params.Encode()
	}

	u := c.baseURL.ResolveReference(rel)

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	return req, nil
}

func (c *RestClient) QueryPrice(ctx context.Context, request PriceRequest) (response PriceResponse, err error) {
	req, err := c.NewRequest(ctx, "GET", "/gateway/v1/price", request.Values())
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response, nil
}
