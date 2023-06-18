package models

import "fmt"

const (
	baseCurrencyID  = "btc-bitcoin"
	quoteCurrencyID = "uah-ukrainian-hryvnia"
	amount          = 1
)

// Converter - struct for build request to thirdpart api
type Converter struct {
	baseCurrencyID  string
	quoteCurrencyID string
	amount          int64
}

func (that *Converter) GetQueryParams() string {
	return fmt.Sprintf("?base_currency_id=%s&quote_currency_id=%s&amount=%d", that.baseCurrencyID, that.quoteCurrencyID, that.amount)
}

func NewConverter() *Converter {
	return &Converter{
		baseCurrencyID:  baseCurrencyID,
		quoteCurrencyID: quoteCurrencyID,
		amount:          amount,
	}
}

// User - struct for create, write, update , delete user
type User struct {
	Email string `toml:"email"`
}

func NewUser(email string) *User {
	return &User{
		Email: email,
	}
}
