package usecase

import (
	"errors"
	"fmt"

	"bitcoin_checker_api/config"
	"bitcoin_checker_api/internal/pkg/email"
	exchange_rate "bitcoin_checker_api/internal/pkg/exchange-rate"
	"bitcoin_checker_api/internal/repository"
	"bitcoin_checker_api/internal/validator"
)

type UseCase struct {
	repository   repository.Repository
	exchangeRate exchange_rate.ExchangeRater
	emailService email.Emailer
}

type UseCaseConfig struct {
	ExchangeRate *config.ExchangeRate
	EmailService *config.EmailService
}

func NewUseCase(c *UseCaseConfig, r repository.Repository) *UseCase {
	return &UseCase{
		repository:   r,
		exchangeRate: exchange_rate.NewExchangeRate(c.ExchangeRate),
		emailService: email.NewEmailService(c.EmailService),
	}
}

func (that *UseCase) SubscribeEmailOnExchangeRate(e string) error {
	validEmail, ok := validator.ValidMailAddress(e)
	if !ok {
		return fmt.Errorf("inalid Email address: %s", e)
	}
	if err := that.repository.Write(validEmail); err != nil {
		return err
	}
	return nil
}

func (that *UseCase) SendEmailsWithExchangeRate() error {
	users := that.repository.ReadAll()
	if len(users) == 0 {
		return errors.New("storage is empty")
	}

	data, err := that.ExchangeRate()
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

func (that *UseCase) ExchangeRate() (string, error) {
	data, err := that.exchangeRate.Get()
	if err != nil {
		return "", err
	}
	return data, nil
}
