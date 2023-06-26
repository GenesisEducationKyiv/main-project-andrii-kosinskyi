package exchange_rate

import (
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

func (that *ExchangeRate) Get() (string, error) {
	res, err := http.Get(that.url)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return "", err
	}
	return string(body), nil
}
