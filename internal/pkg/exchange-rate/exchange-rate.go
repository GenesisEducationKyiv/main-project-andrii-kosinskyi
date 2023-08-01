package exchangerate

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"bitcoin_checker_api/internal/model"
	"bitcoin_checker_api/internal/pkg/mapper"

	"bitcoin_checker_api/config"
)

type ExchangeRate struct {
	serviceName string
	url         string
	next        ExchangeRater
	mapper      mapper.Mapper
}

func NewExchangeRate(c *config.DefaultExchangeRate) (*ExchangeRate, error) {
	erMapper, err := mapper.NewExchangeRateMapper(c.ServiceName)
	if err != nil {
		return nil, err
	}
	return &ExchangeRate{
		serviceName: c.ServiceName,
		url:         fmt.Sprintf(c.URLMask, c.InRate, c.OutRate),
		mapper:      erMapper,
	}, nil
}

func (that *ExchangeRate) SetNext(next ExchangeRater) {
	that.next = next
}

func (that *ExchangeRate) Get(ctx context.Context) (*model.ExchangeRate, error) {
	body, err := that.get(ctx)
	if err != nil {
		next := that.next
		if next == nil {
			return nil, err
		}
		body, err = next.Get(ctx)
	}
	return body, err
}

func (that *ExchangeRate) get(ctx context.Context) (*model.ExchangeRate, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, that.url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ExchageRate.Get() Response: %s status code: %d", body, res.StatusCode)
	}
	er, err := that.mapper.Map(body)
	if err != nil {
		return nil, err
	}
	return er, nil
}
