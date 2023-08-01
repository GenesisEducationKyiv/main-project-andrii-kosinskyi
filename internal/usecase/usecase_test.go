package usecase_test

import (
	"context"
	"errors"
	"testing"

	"bitcoin_checker_api/internal/pkg/mapper"

	"bitcoin_checker_api/internal/usecase"

	"bitcoin_checker_api/config"
	"bitcoin_checker_api/internal/pkg/email"
	exchangerate "bitcoin_checker_api/internal/pkg/exchange-rate"
	"bitcoin_checker_api/internal/repository"
)

func TestUseCase_ExchangeRate(t *testing.T) {
	ctx := context.Background()
	r, _ := repository.NewMockRepository(&config.Storage{})
	ex := exchangerate.NewMockExchangeRate(&config.ExchangeRate{
		URLMask: "mock",
		InRate:  "mock",
		OutRate: "mock",
	})
	m, err := mapper.NewMockExchangeRateMapper("mock")
	if err != nil {
		t.Errorf("TestUseCase_ExchangeRate() err = %v", err)
	}
	es := email.NewMockService(&config.EmailService{})

	useCase := usecase.NewUseCase(r, ex, m, es)
	_, err = useCase.ExchangeRate(ctx)
	if err != nil {
		t.Errorf("TestUseCase_ExchangeRate() err = %v", err)
	}
}

func TestUseCase_ExchangeRateWithError(t *testing.T) {
	ctx := context.Background()
	r, _ := repository.NewMockRepository(&config.Storage{})
	ex := exchangerate.NewMockExchangeRate(&config.ExchangeRate{})
	m, err := mapper.NewMockExchangeRateMapper("mock")
	if err != nil {
		t.Errorf("TestUseCase_ExchangeRate() err = %v", err)
	}
	es := email.NewMockService(&config.EmailService{})

	useCase := usecase.NewUseCase(r, ex, m, es)
	_, err = useCase.ExchangeRate(ctx)
	if err == nil {
		t.Errorf("TestUseCase_ExchangeRateWithError() err = %v", err)
	}
}

func TestUseCase_SubscribeEmailOnExchangeRate(t *testing.T) {
	r, _ := repository.NewMockRepository(&config.Storage{Path: "./storage.json"})
	ex := exchangerate.NewMockExchangeRate(&config.ExchangeRate{})
	m, err := mapper.NewMockExchangeRateMapper("mock")
	if err != nil {
		t.Errorf("TestUseCase_ExchangeRate() err = %v", err)
	}
	es := email.NewMockService(&config.EmailService{})
	useCase := usecase.NewUseCase(r, ex, m, es)

	if err = useCase.SubscribeEmailOnExchangeRate("taras@shchevchenko.com"); err != nil {
		t.Errorf("TestUseCase_SubscribeEmailOnExchangeRate() err = %v", err)
	}
}

func TestUseCase_SubscribeEmailOnExchangeRateWithDuplicateEmailError(t *testing.T) {
	r, _ := repository.NewMockRepository(&config.Storage{})
	ex := exchangerate.NewMockExchangeRate(&config.ExchangeRate{})
	m, err := mapper.NewMockExchangeRateMapper("mock")
	if err != nil {
		t.Errorf("TestUseCase_ExchangeRate() err = %v", err)
	}
	es := email.NewMockService(&config.EmailService{})
	useCase := usecase.NewUseCase(r, ex, m, es)

	err = r.Write("taras@shchevchenko.com")
	if err != nil {
		t.Errorf("TestUseCase_SubscribeEmailOnExchangeRateWithDuplicateEmailError() r.Write() err = %v", err)
	}

	if err = useCase.SubscribeEmailOnExchangeRate("taras@shchevchenko.com"); !errors.Is(err, repository.ErrRecordExists) {
		t.Errorf("TestUseCase_SubscribeEmailOnExchangeRateWithDuplicateEmailError() err = %v", err)
	}
}

func TestUseCase_SendEmailsWithExchangeRate(t *testing.T) {
	ctx := context.Background()
	r, _ := repository.NewMockRepository(&config.Storage{})
	ex := exchangerate.NewMockExchangeRate(&config.ExchangeRate{
		URLMask: "mock",
		InRate:  "mock",
		OutRate: "mock",
	})
	m, err := mapper.NewMockExchangeRateMapper("mock")
	if err != nil {
		t.Errorf("TestUseCase_ExchangeRate() err = %v", err)
	}
	es := email.NewMockService(&config.EmailService{APIKey: "test", FromAddress: "test", FromName: "test"})
	useCase := usecase.NewUseCase(r, ex, m, es)

	err = r.Write("taras@shchevchenko.com")
	if err != nil {
		t.Errorf("TestUseCase_SendEmailsWithExchangeRate() r.Write() err = %v", err)
	}

	if err = useCase.SendEmailsWithExchangeRate(ctx); err != nil {
		t.Errorf("TestUseCase_SendEmailsWithExchangeRate() err = %v", err)
	}
}

func TestUseCase_SendEmailsWithExchangeRateWithErrorEmptyRepository(t *testing.T) {
	ctx := context.Background()
	r, _ := repository.NewMockRepository(&config.Storage{})
	ex := exchangerate.NewMockExchangeRate(&config.ExchangeRate{})
	m, err := mapper.NewMockExchangeRateMapper("mock")
	if err != nil {
		t.Errorf("TestUseCase_ExchangeRate() err = %v", err)
	}
	es := email.NewMockService(&config.EmailService{APIKey: "", FromAddress: "", FromName: ""})
	useCase := usecase.NewUseCase(r, ex, m, es)

	if err = useCase.SendEmailsWithExchangeRate(ctx); !errors.Is(err, usecase.ErrUseCaseEmptyUserList) {
		t.Errorf("TestUseCase_SendEmailsWithExchangeRate() err = %v", err)
	}
}

func TestUseCase_SendEmailsWithExchangeRateWithErrorInRequestToExchangeRate(t *testing.T) {
	ctx := context.Background()
	r, _ := repository.NewMockRepository(&config.Storage{})
	ex := exchangerate.NewMockExchangeRate(&config.ExchangeRate{})
	m, err := mapper.NewMockExchangeRateMapper("mock")
	if err != nil {
		t.Errorf("TestUseCase_ExchangeRate() err = %v", err)
	}
	es := email.NewMockService(&config.EmailService{APIKey: "", FromAddress: "", FromName: ""})
	useCase := usecase.NewUseCase(r, ex, m, es)

	err = r.Write("taras@shchevchenko.com")
	if err != nil {
		t.Errorf("TestUseCase_SendEmailsWithExchangeRate() r.Write() err = %v", err)
	}

	if err = useCase.SendEmailsWithExchangeRate(ctx); err == nil {
		t.Errorf("TestUseCase_SendEmailsWithExchangeRate() err = %v", err)
	}
}

func TestUseCase_SendEmailsWithExchangeRateWithErrorSendEmail(t *testing.T) {
	ctx := context.Background()
	r, _ := repository.NewMockRepository(&config.Storage{})
	ex := exchangerate.NewMockExchangeRate(&config.ExchangeRate{})
	m, err := mapper.NewMockExchangeRateMapper("mock")
	if err != nil {
		t.Errorf("TestUseCase_ExchangeRate() err = %v", err)
	}
	es := email.NewMockService(&config.EmailService{APIKey: "", FromAddress: "", FromName: ""})
	useCase := usecase.NewUseCase(r, ex, m, es)

	err = r.Write("taras@shchevchenko.com")
	if err != nil {
		t.Errorf("TestUseCase_SendEmailsWithExchangeRate() r.Write() err = %v", err)
	}

	if err = useCase.SendEmailsWithExchangeRate(ctx); err == nil {
		t.Errorf("TestUseCase_SendEmailsWithExchangeRate() err = %v", err)
	}
}
