package model

import "time"

type ExchangeRate struct {
	BaseCurrency  string
	QuoteCurrency string
	Price         float64
	Date          time.Time
}

func NewExchangeRate(baseCurrency, quoteCurrency string, price float64) *ExchangeRate {
	return &ExchangeRate{
		BaseCurrency:  baseCurrency,
		QuoteCurrency: quoteCurrency,
		Price:         price,
		Date:          time.Now(),
	}
}
