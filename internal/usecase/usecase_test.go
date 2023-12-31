package usecase_test

import (
	"context"
	"errors"
	"testing"

	"bitcoin_checker_api/internal/usecase"

	"bitcoin_checker_api/config"
	"bitcoin_checker_api/internal/pkg/email"
	exchangerate "bitcoin_checker_api/internal/pkg/exchange-rate"
	"bitcoin_checker_api/internal/repository"
)

func TestUseCase_ExchangeRate(t *testing.T) {
	ctx := context.Background()
	r, _ := repository.NewMockRepository(&config.Storage{})
	excRateCoinpaprika, err := exchangerate.NewMockExchangeRate(&config.DefaultExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "btc-bitcoin",
		OutRate: "uah-ukrainian-hryvnia",
	})
	if err != nil {
		t.Errorf("TestUseCase_ExchangeRate() err = %v", err)
	}
	excRateBinance, err := exchangerate.NewMockExchangeRate(&config.DefaultExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "btc-bitcoin",
		OutRate: "uah-ukrainian-hryvnia",
	})
	if err != nil {
		t.Errorf("TestUseCase_ExchangeRate() err = %v", err)
	}
	excRateCoinpaprika.SetNext(excRateBinance)
	es := email.NewMockService(&config.EmailService{})

	useCase := usecase.NewUseCase(r, excRateCoinpaprika, es)
	_, err = useCase.ExchangeRate(ctx)
	if err != nil {
		t.Errorf("TestUseCase_ExchangeRate() err = %v", err)
	}
}

func TestUseCase_ExchangeRateWithError(t *testing.T) {
	ctx := context.Background()
	r, _ := repository.NewMockRepository(&config.Storage{})
	excRateCoinpaprika, err := exchangerate.NewMockExchangeRate(&config.DefaultExchangeRate{})
	if err != nil {
		t.Errorf("TestUseCase_ExchangeRate() err = %v", err)
	}
	excRateBinance, err := exchangerate.NewMockExchangeRate(&config.DefaultExchangeRate{})
	if err != nil {
		t.Errorf("TestUseCase_ExchangeRate() err = %v", err)
	}
	excRateCoinpaprika.SetNext(excRateBinance)
	es := email.NewMockService(&config.EmailService{})

	useCase := usecase.NewUseCase(r, excRateCoinpaprika, es)
	_, err = useCase.ExchangeRate(ctx)
	if err == nil {
		t.Errorf("TestUseCase_ExchangeRateWithError() err = %v", err)
	}
}

func TestUseCase_SubscribeEmailOnExchangeRate(t *testing.T) {
	r, _ := repository.NewMockRepository(&config.Storage{Path: "./storage.json"})
	excRateCoinpaprika, err := exchangerate.NewMockExchangeRate(&config.DefaultExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "btc-bitcoin",
		OutRate: "uah-ukrainian-hryvnia",
	})
	if err != nil {
		t.Errorf("TestUseCase_ExchangeRate() err = %v", err)
	}
	excRateBinance, err := exchangerate.NewMockExchangeRate(&config.DefaultExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "btc-bitcoin",
		OutRate: "uah-ukrainian-hryvnia",
	})
	if err != nil {
		t.Errorf("TestUseCase_ExchangeRate() err = %v", err)
	}
	excRateCoinpaprika.SetNext(excRateBinance)
	es := email.NewMockService(&config.EmailService{})
	useCase := usecase.NewUseCase(r, excRateCoinpaprika, es)

	if err = useCase.SubscribeEmailOnExchangeRate("taras@shchevchenko.com"); err != nil {
		t.Errorf("TestUseCase_SubscribeEmailOnExchangeRate() err = %v", err)
	}
}

func TestUseCase_SubscribeEmailOnExchangeRateWithDuplicateEmailError(t *testing.T) {
	r, _ := repository.NewMockRepository(&config.Storage{})
	excRateCoinpaprika, err := exchangerate.NewMockExchangeRate(&config.DefaultExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "btc-bitcoin",
		OutRate: "uah-ukrainian-hryvnia",
	})
	if err != nil {
		t.Errorf("TestUseCase_ExchangeRate() err = %v", err)
	}
	excRateBinance, err := exchangerate.NewMockExchangeRate(&config.DefaultExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "btc-bitcoin",
		OutRate: "uah-ukrainian-hryvnia",
	})
	if err != nil {
		t.Errorf("TestUseCase_ExchangeRate() err = %v", err)
	}
	excRateCoinpaprika.SetNext(excRateBinance)
	es := email.NewMockService(&config.EmailService{})
	useCase := usecase.NewUseCase(r, excRateCoinpaprika, es)

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
	excRateCoinpaprika, err := exchangerate.NewMockExchangeRate(&config.DefaultExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "btc-bitcoin",
		OutRate: "uah-ukrainian-hryvnia",
	})
	if err != nil {
		t.Errorf("TestUseCase_ExchangeRate() err = %v", err)
	}
	excRateBinance, err := exchangerate.NewMockExchangeRate(&config.DefaultExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "btc-bitcoin",
		OutRate: "uah-ukrainian-hryvnia",
	})
	if err != nil {
		t.Errorf("TestUseCase_ExchangeRate() err = %v", err)
	}
	excRateCoinpaprika.SetNext(excRateBinance)
	es := email.NewMockService(&config.EmailService{APIKey: "test", FromAddress: "test", FromName: "test"})
	useCase := usecase.NewUseCase(r, excRateCoinpaprika, es)

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
	excRateCoinpaprika, err := exchangerate.NewMockExchangeRate(&config.DefaultExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "btc-bitcoin",
		OutRate: "uah-ukrainian-hryvnia",
	})
	if err != nil {
		t.Errorf("TestUseCase_ExchangeRate() err = %v", err)
	}
	excRateBinance, err := exchangerate.NewMockExchangeRate(&config.DefaultExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "btc-bitcoin",
		OutRate: "uah-ukrainian-hryvnia",
	})
	if err != nil {
		t.Errorf("TestUseCase_ExchangeRate() err = %v", err)
	}
	excRateCoinpaprika.SetNext(excRateBinance)
	es := email.NewMockService(&config.EmailService{APIKey: "", FromAddress: "", FromName: ""})
	useCase := usecase.NewUseCase(r, excRateCoinpaprika, es)

	if err = useCase.SendEmailsWithExchangeRate(ctx); !errors.Is(err, usecase.ErrUseCaseEmptyUserList) {
		t.Errorf("TestUseCase_SendEmailsWithExchangeRate() err = %v", err)
	}
}

func TestUseCase_SendEmailsWithExchangeRateWithErrorInRequestToExchangeRate(t *testing.T) {
	ctx := context.Background()
	r, _ := repository.NewMockRepository(&config.Storage{})
	excRateCoinpaprika, err := exchangerate.NewMockExchangeRate(&config.DefaultExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "btc-bitcoin",
		OutRate: "uah-ukrainian-hryvnia",
	})
	if err != nil {
		t.Errorf("TestUseCase_ExchangeRate() err = %v", err)
	}
	excRateBinance, err := exchangerate.NewMockExchangeRate(&config.DefaultExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "btc-bitcoin",
		OutRate: "uah-ukrainian-hryvnia",
	})
	if err != nil {
		t.Errorf("TestUseCase_ExchangeRate() err = %v", err)
	}
	excRateCoinpaprika.SetNext(excRateBinance)
	es := email.NewMockService(&config.EmailService{APIKey: "", FromAddress: "", FromName: ""})
	useCase := usecase.NewUseCase(r, excRateCoinpaprika, es)

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
	excRateCoinpaprika, err := exchangerate.NewMockExchangeRate(&config.DefaultExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "btc-bitcoin",
		OutRate: "uah-ukrainian-hryvnia",
	})
	if err != nil {
		t.Errorf("TestUseCase_ExchangeRate() err = %v", err)
	}
	excRateBinance, err := exchangerate.NewMockExchangeRate(&config.DefaultExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "btc-bitcoin",
		OutRate: "uah-ukrainian-hryvnia",
	})
	if err != nil {
		t.Errorf("TestUseCase_ExchangeRate() err = %v", err)
	}
	excRateCoinpaprika.SetNext(excRateBinance)
	es := email.NewMockService(&config.EmailService{APIKey: "", FromAddress: "", FromName: ""})
	useCase := usecase.NewUseCase(r, excRateCoinpaprika, es)

	err = r.Write("taras@shchevchenko.com")
	if err != nil {
		t.Errorf("TestUseCase_SendEmailsWithExchangeRate() r.Write() err = %v", err)
	}

	if err = useCase.SendEmailsWithExchangeRate(ctx); err == nil {
		t.Errorf("TestUseCase_SendEmailsWithExchangeRate() err = %v", err)
	}
}
