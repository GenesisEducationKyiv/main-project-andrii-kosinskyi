package usecase

import (
	"context"
	"fmt"

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

func NewUseCase(r repository.Repository, er exchangerate.ExchangeRater, m mapper.Mapper, es email.Emailer) *UseCase {
	return &UseCase{
		repository:   r,
		exchangeRate: er,
		mapper:       m,
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
	data, err := that.exchangeRate.Get(ctx)
	if err != nil {
		return nil, err
	}

	exchangeRate, err := that.mapper.Map(data)
	if err != nil {
		return nil, fmt.Errorf("error from %s : %w", that.mapper.Name(), err)
	}

	return exchangeRate, nil
}
