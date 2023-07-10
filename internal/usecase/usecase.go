package usecase

import (
	"context"

	"bitcoin_checker_api/internal/model"
	"bitcoin_checker_api/internal/pkg/mapper"

	"bitcoin_checker_api/internal/pkg/email"
	exchangerate "bitcoin_checker_api/internal/pkg/exchange-rate"
	"bitcoin_checker_api/internal/repository"
)

type UseCase struct {
	repository   repository.Repository
	exchangeRate exchangerate.ExchangeRater
	mapper       mapper.Mapper
	emailService email.Emailer
}

func NewUseCase(r repository.Repository, er exchangerate.ExchangeRater, es email.Emailer) *UseCase {
	return &UseCase{
		repository:   r,
		exchangeRate: er,
		emailService: es,
	}
}

func (that *UseCase) SubscribeEmailOnExchangeRate(e string) error {
	return that.repository.Write(e)
}

func (that *UseCase) SendEmailsWithExchangeRate(ctx context.Context) error {
	users := that.repository.ReadAll()
	if len(users) == 0 {
		return ErrUseCaseEmptyUserList
	}

	exchangeRate, err := that.ExchangeRate(ctx)
	if err != nil {
		return err
	}

	for _, user := range users {
		if err = that.emailService.Send(user.Email, exchangeRate); err != nil {
			return err
		}
	}
	return nil
}

func (that *UseCase) ExchangeRate(ctx context.Context) (*model.ExchangeRate, error) {
	exchangeRate, err := that.exchangeRate.Get(ctx)
	if err != nil {
		return nil, err
	}

	return exchangeRate, nil
}
