package mapper

import (
	"testing"
)

func TestName(t *testing.T) {
	if service := NewCoinPaprikaMapper(); service.Name() != CoinPaprikaService {
		t.Errorf("TestName(): want: %s got: %s ", CoinPaprikaService, service.Name())
	}
}

func TestCoinPaprikaMapper_Map(t *testing.T) {
	validRespBody := []byte("{\n\"base_currency_id\": \"btc-bitcoin\",\n\"base_currency_name\": \"Bitcoin\",\n\"base_price_last_updated\": \"2023-07-07T20:09:17Z\",\n\"quote_currency_id\": \"uah-ukrainian-hryvnia\",\n\"quote_currency_name\": \"Ukrainian Hryvnia\",\n\"quote_price_last_updated\": \"2023-07-07T19:43:37Z\",\n\"amount\": 1,\n\"price\": 1117038.070894138\n}")
	if _, err := NewCoinPaprikaMapper().Map(validRespBody); err != nil {
		t.Errorf("TestCoinPaprikaMapper_MapWithError(): err: %v", err)
	}
}

func TestCoinPaprikaMapper_MapWithError(t *testing.T) {
	if _, err := NewCoinPaprikaMapper().Map(make([]byte, 0)); err != ErrUnmarshal {
		t.Errorf("TestCoinPaprikaMapper_MapWithError(): want: %v got: %v", ErrUnmarshal, err)
	}
}
