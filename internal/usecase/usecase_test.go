package usecase

import (
	"context"
	"testing"

	"bitcoin_checker_api/config"
	"bitcoin_checker_api/internal/pkg/email"
	exchangerate "bitcoin_checker_api/internal/pkg/exchange-rate"
	"bitcoin_checker_api/internal/repository"
)

func TestUseCase_ExchangeRate(t *testing.T) {
	ctx := context.Background()
	r, _ := repository.NewMockRepository(&config.Storage{})
	cfg := &config.ExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "btc-bitcoin",
		OutRate: "uah-ukrainian-hryvnia",
	}
	ex := exchangerate.NewExchangeRate(cfg)
	es := email.NewMockService(&config.EmailService{})

	useCase := NewUseCase(r, ex, es)
	_, err := useCase.ExchangeRate(ctx)
	if err != nil {
		t.Errorf("TestUseCase_ExchangeRate() err = %v", err)
	}
}

func TestUseCase_ExchangeRateWithError(t *testing.T) {
	ctx := context.Background()
	r, _ := repository.NewMockRepository(&config.Storage{})
	cfg := &config.ExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "",
		OutRate: "uah-ukrainian-hryvnia",
	}
	ex := exchangerate.NewExchangeRate(cfg)
	es := email.NewMockService(&config.EmailService{})

	useCase := NewUseCase(r, ex, es)
	_, err := useCase.ExchangeRate(ctx)
	if err == nil {
		t.Errorf("TestUseCase_ExchangeRateWithError() err = %v", err)
	}
}

func TestUseCase_SubscribeEmailOnExchangeRate(t *testing.T) {
	r, _ := repository.NewMockRepository(&config.Storage{Path: "./storage.json"})
	ex := exchangerate.NewExchangeRate(&config.ExchangeRate{})
	es := email.NewMockService(&config.EmailService{})
	useCase := NewUseCase(r, ex, es)

	if err := useCase.SubscribeEmailOnExchangeRate("taras@shchevchenko.com"); err != nil {
		t.Errorf("TestUseCase_SubscribeEmailOnExchangeRate() err = %v", err)
	}
}

func TestUseCase_SubscribeEmailOnExchangeRateWithDuplicateEmailError(t *testing.T) {
	r, _ := repository.NewMockRepository(&config.Storage{})
	ex := exchangerate.NewExchangeRate(&config.ExchangeRate{})
	es := email.NewMockService(&config.EmailService{})
	useCase := NewUseCase(r, ex, es)

	err := r.Write("taras@shchevchenko.com")
	if err != nil {
		t.Errorf("TestUseCase_SubscribeEmailOnExchangeRateWithDuplicateEmailError() r.Write() err = %v", err)
	}

	if err = useCase.SubscribeEmailOnExchangeRate("taras@shchevchenko.com"); err != repository.ErrRecordExists {
		t.Errorf("TestUseCase_SubscribeEmailOnExchangeRateWithDuplicateEmailError() err = %v", err)
	}
}

func TestUseCase_SendEmailsWithExchangeRate(t *testing.T) {
	ctx := context.Background()
	r, _ := repository.NewMockRepository(&config.Storage{})
	cfg := &config.ExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "btc-bitcoin",
		OutRate: "uah-ukrainian-hryvnia",
	}
	ex := exchangerate.NewExchangeRate(cfg)
	es := email.NewMockService(&config.EmailService{APIKey: "test", FromAddress: "test", FromName: "test"})
	useCase := NewUseCase(r, ex, es)

	err := r.Write("taras@shchevchenko.com")
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
	cfg := &config.ExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "btc-bitcoin",
		OutRate: "uah-ukrainian-hryvnia",
	}
	ex := exchangerate.NewExchangeRate(cfg)
	es := email.NewMockService(&config.EmailService{APIKey: "", FromAddress: "", FromName: ""})
	useCase := NewUseCase(r, ex, es)

	if err := useCase.SendEmailsWithExchangeRate(ctx); err != ErrUseCaseEmptyUserList {
		t.Errorf("TestUseCase_SendEmailsWithExchangeRate() err = %v", err)
	}
}

func TestUseCase_SendEmailsWithExchangeRateWithErrorInRequestToExchangeRate(t *testing.T) {
	ctx := context.Background()
	r, _ := repository.NewMockRepository(&config.Storage{})
	cfg := &config.ExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "",
		OutRate: "uah-ukrainian-hryvnia",
	}
	ex := exchangerate.NewExchangeRate(cfg)
	es := email.NewMockService(&config.EmailService{APIKey: "", FromAddress: "", FromName: ""})
	useCase := NewUseCase(r, ex, es)

	err := r.Write("taras@shchevchenko.com")
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
	cfg := &config.ExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "btc-bitcoin",
		OutRate: "uah-ukrainian-hryvnia",
	}
	ex := exchangerate.NewExchangeRate(cfg)
	es := email.NewMockService(&config.EmailService{APIKey: "", FromAddress: "", FromName: ""})
	useCase := NewUseCase(r, ex, es)

	err := r.Write("taras@shchevchenko.com")
	if err != nil {
		t.Errorf("TestUseCase_SendEmailsWithExchangeRate() r.Write() err = %v", err)
	}

	if err = useCase.SendEmailsWithExchangeRate(ctx); err == nil {
		t.Errorf("TestUseCase_SendEmailsWithExchangeRate() err = %v", err)
	}
}
