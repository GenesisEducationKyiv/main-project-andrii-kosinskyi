package exchangerate_test

import (
	"context"
	"testing"

	exchangerate "bitcoin_checker_api/internal/pkg/exchange-rate"

	"bitcoin_checker_api/config"
)

func TestExchangeRate_Get(t *testing.T) {
	ctx := context.Background()
	client := exchangerate.NewExchangeRate(&config.ExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "btc-bitcoin",
		OutRate: "uah-ukrainian-hryvnia",
	})

	_, err := client.Get(ctx)
	if err != nil {
		t.Errorf("TestExchangeRate_Get() err = %v", err)
	}
}
