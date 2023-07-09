package mapper

import (
	"encoding/json"
	"time"

	"bitcoin_checker_api/internal/model"
)

type CoinPaprikaMapper struct {
	name string
}

type CoinPaprikaRespBody struct {
	BaseCurrencyId        string    `json:"base_currency_id"`
	BaseCurrencyName      string    `json:"base_currency_name"`
	BasePriceLastUpdated  time.Time `json:"base_price_last_updated"`
	QuoteCurrencyId       string    `json:"quote_currency_id"`
	QuoteCurrencyName     string    `json:"quote_currency_name"`
	QuotePriceLastUpdated time.Time `json:"quote_price_last_updated"`
	Amount                int       `json:"amount"`
	Price                 float64   `json:"price"`
}

func NewCoinPaprikaMapper() Mapper {
	return &CoinPaprikaMapper{
		name: CoinPaprikaService,
	}
}

func (that *CoinPaprikaMapper) Name() string {
	return that.name
}

func (that *CoinPaprikaMapper) Map(serviceRespBody []byte) (*model.ExchangeRate, error) {
	cpm := &CoinPaprikaRespBody{}
	if err := json.Unmarshal(serviceRespBody, cpm); err != nil {
		return nil, ErrUnmarshal
	}
	return &model.ExchangeRate{
		BaseCurrency:  cpm.BaseCurrencyName,
		QuoteCurrency: cpm.QuoteCurrencyName,
		Price:         cpm.Price,
		Date:          time.Now(),
	}, nil
}
