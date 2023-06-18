package models

import "fmt"

const (
	baseCurrencyID  = "btc-bitcoin"
	quoteCurrencyID = "uah-ukrainian-hryvnia"
	amount          = 1
)

// Converter - struct for build request to third-part api.
type Converter struct {
	baseCurrencyID  string
	quoteCurrencyID string
	amount          int64
}

func (that *Converter) GetQueryParams() string {
	format := "?base_currency_id=%s&quote_currency_id=%s&amount=%d"
	return fmt.Sprintf(format, that.baseCurrencyID, that.quoteCurrencyID, that.amount)
}

func NewConverter() *Converter {
	return &Converter{
		baseCurrencyID:  baseCurrencyID,
		quoteCurrencyID: quoteCurrencyID,
		amount:          amount,
	}
}

// User - struct for create, write, update , delete user.
type User struct {
	Email string `toml:"email"`
}

func NewUser(email string) *User {
	return &User{
		Email: email,
	}
}
