package exchangerate

import (
	"context"
	"fmt"
	"net/http"

	"bitcoin_checker_api/config"
)

type MockExchangeRate struct {
	url string
}

func NewMockExchangeRate(c *config.DefaultExchangeRate) *MockExchangeRate {
	return &MockExchangeRate{url: fmt.Sprintf("%s%s%s", c.URLMask, c.InRate, c.OutRate)}
}

func (that *MockExchangeRate) Get(_ context.Context) ([]byte, error) {
	if that.url == "" {
		return make([]byte, 0), fmt.Errorf("ExchageRate.Get() Response: %s status code: %d", "", http.StatusNotFound)
	}

	return make([]byte, 1), nil
}
