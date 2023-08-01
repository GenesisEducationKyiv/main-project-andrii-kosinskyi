package mapper

import (
	"errors"

	"bitcoin_checker_api/internal/model"
)

func NewMockExchangeRateMapper(serviceName string) (Mapper, error) {
	switch {
	case serviceName != "":
		return NewMockMapper(), nil
	default:
		return nil, ErrUnknownService
	}
}

type MockMapper struct {
	name string
}

func NewMockMapper() *MockMapper {
	return &MockMapper{name: "mock mapper"}
}

func (that *MockMapper) Name() string {
	return that.name
}

func (that *MockMapper) Map(serviceRespBody []byte) (*model.ExchangeRate, error) {
	if len(serviceRespBody) == 0 {
		return nil, errors.New("marshaling error")
	}
	return &model.ExchangeRate{}, nil
}
