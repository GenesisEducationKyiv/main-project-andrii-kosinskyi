package exchangerate_test

import (
	"context"
	"testing"

	exchangerate "bitcoin_checker_api/internal/pkg/exchange-rate"

	"bitcoin_checker_api/config"
)

func TestExchangeRate_Get_Coinpaprika(t *testing.T) {
	ctx := context.Background()
	client, _ := exchangerate.NewExchangeRate(&config.DefaultExchangeRate{
		ServiceName: "coinpaprika",
		URLMask:     "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:      "btc-bitcoin",
		OutRate:     "uah-ukrainian-hryvnia",
	})

	_, err := client.Get(ctx)
	if err != nil {
		t.Errorf("TestExchangeRate_Get_Coinpaprika() err = %v", err)
	}
}

//func TestExchangeRate_Get_Binance(t *testing.T) {
//	ctx := context.Background()
//	client, _ := exchangerate.NewExchangeRate(&config.DefaultExchangeRate{
//		ServiceName: "binance",
//		URLMask:     "https://api.binance.com/api/v3/ticker/price?symbol=%s%s",
//		InRate:      "BTC",
//		OutRate:     "UAH",
//	})
//
//	_, err := client.Get(ctx)
//	if err != nil {
//		t.Errorf("TestExchangeRate_Get_Binance() err = %v", err)
//	}
//}

func TestExchangeRate_SetNext(t *testing.T) {
	ctx := context.Background()
	service1, _ := exchangerate.NewExchangeRate(&config.DefaultExchangeRate{
		ServiceName: "coinpaprika",
		URLMask:     "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:      "btc-bitcoin",
		OutRate:     "uah-ukrainian-hryvnia",
	})
	service2, _ := exchangerate.NewExchangeRate(&config.DefaultExchangeRate{
		ServiceName: "binance",
		URLMask:     "https://api.binance.com/api/v3/ticker/price?symbol=%s%s",
		InRate:      "BTC",
		OutRate:     "UAH",
	})
	service1.SetNext(service2)
	_, err := service1.Get(ctx)
	if err != nil {
		t.Errorf("TestExchangeRate_Get() err = %v", err)
	}
}
