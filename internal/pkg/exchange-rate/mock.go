package exchangerate

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"bitcoin_checker_api/config"
)

type MockExchangeRate struct {
	url string
}

func NewMockExchangeRate(c *config.ExchangeRate) *MockExchangeRate {
	return &MockExchangeRate{url: fmt.Sprintf(c.URLMask, c.InRate, c.OutRate)}
}

func (that *MockExchangeRate) Get(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, that.url, nil)
	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("ExchageRate.Get() Response: %s status code: %d", body, res.StatusCode)
	}

	return string(body), nil
}
