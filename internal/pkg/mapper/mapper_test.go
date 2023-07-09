package mapper

import "testing"

func TestNewExchangeRateMapper(t *testing.T) {
	tests := []struct {
		name        string
		serviceName string
		want        string
	}{
		{name: "Create CoinPaprika mapper", serviceName: CoinPaprikaService, want: CoinPaprikaService},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if service, err := NewExchangeRateMapper(tt.serviceName); err != nil && service.Name() != tt.want {
				t.Errorf("")
			}
		})
	}
}

func TestNewExchangeRateMapperWithError(t *testing.T) {
	if _, err := NewExchangeRateMapper("Unknown service name"); err != ErrUnknownService {
		t.Errorf("TestNewExchangeRateMapperWithError() want: %v got: %v", ErrUnknownService, err)
	}
}
