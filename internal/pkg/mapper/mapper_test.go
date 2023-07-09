package mapper_test

import (
	"bitcoin_checker_api/internal/pkg/mapper"
	"errors"
	"testing"
)

func TestNewExchangeRateMapper(t *testing.T) {
	tests := []struct {
		name        string
		serviceName string
		want        string
	}{
		{name: "Create CoinPaprika mapper", serviceName: mapper.CoinPaprikaService, want: mapper.CoinPaprikaService},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if service, err := mapper.NewExchangeRateMapper(tt.serviceName); err != nil && service.Name() != tt.want {
				t.Errorf("")
			}
		})
	}
}

func TestNewExchangeRateMapperWithError(t *testing.T) {
	if _, err := mapper.NewExchangeRateMapper("Unknown service name"); !errors.Is(err, mapper.ErrUnknownService) {
		t.Errorf("TestNewExchangeRateMapperWithError() want: %v got: %v", mapper.ErrUnknownService, err)
	}
}
