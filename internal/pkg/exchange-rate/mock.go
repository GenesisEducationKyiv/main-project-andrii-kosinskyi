package exchangerate

import (
	"bitcoin_checker_api/internal/model"
	"context"
	"fmt"
	"net/http"

	"bitcoin_checker_api/config"
)

type MockExchangeRate struct {
	url string
}

func NewMockExchangeRate(c *config.DefaultExchangeRate) (*MockExchangeRate, error) {
	return &MockExchangeRate{url: fmt.Sprintf("%s%s%s", c.URLMask, c.InRate, c.OutRate)}, nil
}

func (that *MockExchangeRate) SetNext(_ ExchangeRater) {

}

func (that *MockExchangeRate) Get(_ context.Context) (*model.ExchangeRate, error) {
	if that.url == "" {
		return nil, fmt.Errorf("ExchageRate.Get() Response: %s status code: %d", "", http.StatusNotFound)
	}

	return &model.ExchangeRate{}, nil
}
