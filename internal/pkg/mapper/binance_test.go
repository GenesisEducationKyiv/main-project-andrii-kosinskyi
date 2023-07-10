package mapper_test

import (
	"bitcoin_checker_api/internal/pkg/mapper"
	"errors"
	"testing"
)

func TestBinanceMapper_Name(t *testing.T) {
	if service := mapper.NewBinanceMapper(); service.Name() != mapper.BinanceService {
		t.Errorf("TestName(): want: %s got: %s ", mapper.BinanceService, service.Name())
	}
}

func TestBinanceMapper_Map(t *testing.T) {
	//nolint:lll
	validRespBody := []byte("{\n\"symbol\": \"BTCUAH\",\n\"price\": \"1162870.00000000\"\n}")
	if _, err := mapper.NewBinanceMapper().Map(validRespBody); err != nil {
		t.Errorf("TestBinanceMapper_Map(): err: %v", err)
	}
}

func TestBinanceMapper_MapWithError(t *testing.T) {
	if _, err := mapper.NewCoinPaprikaMapper().Map(make([]byte, 0)); !errors.Is(err, mapper.ErrUnmarshal) {
		t.Errorf("TestBinanceMapper_MapWithError(): want: %v got: %v", mapper.ErrUnmarshal, err)
	}
}
