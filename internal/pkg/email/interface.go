package email

import "bitcoin_checker_api/internal/model"

type Emailer interface {
	Send(email string, exchangeRate *model.ExchangeRate) error
}
