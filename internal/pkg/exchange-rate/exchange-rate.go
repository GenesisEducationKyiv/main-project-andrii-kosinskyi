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

func (that *ExchangeRate) Get(ctx context.Context) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, that.url, nil)
	if err != nil {
		return make([]byte, 0), err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return make([]byte, 0), err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return make([]byte, 0), err
	}

	if res.StatusCode != http.StatusOK {
		return make([]byte, 0), fmt.Errorf("ExchageRate.Get() Response: %s status code: %d", body, res.StatusCode)
	}

	return body, nil
}
