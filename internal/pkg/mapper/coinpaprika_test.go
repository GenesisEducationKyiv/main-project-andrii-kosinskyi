package mapper_test

import (
	"errors"
	"testing"

	"bitcoin_checker_api/internal/pkg/mapper"
)

func TestCoinPaprikaMapper_Name(t *testing.T) {
	if service := mapper.NewCoinPaprikaMapper(); service.Name() != mapper.CoinPaprikaService {
		t.Errorf("TestName(): want: %s got: %s ", mapper.CoinPaprikaService, service.Name())
	}
}

func TestCoinPaprikaMapper_Map(t *testing.T) {
	//nolint:lll
	validRespBody := []byte("{\n\"base_currency_id\": \"btc-bitcoin\",\n\"base_currency_name\": \"Bitcoin\",\n\"base_price_last_updated\": \"2023-07-07T20:09:17Z\",\n\"quote_currency_id\": \"uah-ukrainian-hryvnia\",\n\"quote_currency_name\": \"Ukrainian Hryvnia\",\n\"quote_price_last_updated\": \"2023-07-07T19:43:37Z\",\n\"amount\": 1,\n\"price\": 1117038.070894138\n}")
	if _, err := mapper.NewCoinPaprikaMapper().Map(validRespBody); err != nil {
		t.Errorf("TestCoinPaprikaMapper_Map(): err: %v", err)
	}
}

func TestCoinPaprikaMapper_MapWithError(t *testing.T) {
	if _, err := mapper.NewCoinPaprikaMapper().Map(make([]byte, 0)); !errors.Is(err, mapper.ErrUnmarshal) {
		t.Errorf("TestCoinPaprikaMapper_MapWithError(): want: %v got: %v", mapper.ErrUnmarshal, err)
	}
}
