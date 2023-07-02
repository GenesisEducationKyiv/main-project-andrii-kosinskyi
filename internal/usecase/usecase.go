package usecase

import (
	"bitcoin_checker_api/config"
	"bitcoin_checker_api/internal/pkg/email"
	exchangerate "bitcoin_checker_api/internal/pkg/exchange-rate"
	"bitcoin_checker_api/internal/repository"
	"context"
	"errors"
)

type UseCase struct {
	repository   repository.Repository
	exchangeRate exchangerate.ExchangeRater
	emailService email.Emailer
}

type Config struct {
	ExchangeRate *config.ExchangeRate
	EmailService *config.EmailService
}

func NewUseCase(c *Config, r repository.Repository) *UseCase {
	return &UseCase{
		repository:   r,
		exchangeRate: exchangerate.NewExchangeRate(c.ExchangeRate),
		emailService: email.NewService(c.EmailService),
	}
}

func (that *UseCase) SubscribeEmailOnExchangeRate(e string) error {
	return that.repository.Write(e)
}

func (that *UseCase) SendEmailsWithExchangeRate(ctx context.Context) error {
	users := that.repository.ReadAll()
	if len(users) == 0 {
		return errors.New("storage is empty")
	}

	data, err := that.ExchangeRate(ctx)
	if err != nil {
		return err
	}

	for _, user := range users {
		if err = that.emailService.Send(user.Email, data); err != nil {
			return err
		}
	}
	return nil
}

func (that *UseCase) ExchangeRate(ctx context.Context) (string, error) {
	data, err := that.exchangeRate.Get(ctx)
	if err != nil {
		return "", err
	}
	return data, nil
}
