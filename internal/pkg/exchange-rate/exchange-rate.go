package exchangerate

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"bitcoin_checker_api/config"
)

type ExchangeRate struct {
	url string
}

func NewExchangeRate(c *config.ExchangeRate) *ExchangeRate {
	return &ExchangeRate{url: fmt.Sprintf(c.URLMask, c.InRate, c.OutRate)}
}

func (that *ExchangeRate) Get(ctx context.Context) (string, error) {
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

	return string(body), nil
}
