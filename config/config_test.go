package config_test

import (
	"errors"
	"testing"

	"bitcoin_checker_api/config"
)

func TestEmailService_Empty(t *testing.T) {
	tests := []struct {
		name string
		cfg  *config.EmailService
		want error
	}{
		{name: "Valid config", cfg: &config.EmailService{
			APIKey:      "qwer-3423d-sdfasdf",
			FromAddress: "tarsa@schecvcnko.com",
			FromName:    "tarsa@schecvcnko.com",
		}, want: nil},
		{name: "Invalid config with empty APIKey", cfg: &config.EmailService{
			APIKey:      "",
			FromAddress: "tarsa@schecvcnko.com",
			FromName:    "tarsa@schecvcnko.com",
		}, want: config.ErrEmptyConfig},
		{name: "Invalid config with empty FromAddress", cfg: &config.EmailService{
			APIKey:      "qwer-3423d-sdfasdf",
			FromAddress: "",
			FromName:    "tarsa@schecvcnko.com",
		}, want: config.ErrEmptyConfig},
		{name: "Invalid config with empty FromName", cfg: &config.EmailService{
			APIKey:      "qwer-3423d-sdfasdf",
			FromAddress: "tarsa@schecvcnko.com",
			FromName:    "",
		}, want: config.ErrEmptyConfig},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cfg.Empty(); !errors.Is(err, tt.want) {
				t.Errorf("TestEmailService_Empty() name = %s err = %v want = %v", tt.name, err, tt.want)
			}
		})
	}
}

func TestExchangeRate_Empty(t *testing.T) {
	tests := []struct {
		name string
		cfg  *config.ExchangeRate
		want error
	}{
		{name: "Valid config", cfg: &config.ExchangeRate{
			Binance: &config.DefaultExchangeRate{
				ServiceName: "bitcoin",
				URLMask:     "https://qwe.wer",
				InRate:      "ua",
				OutRate:     "bit",
			},
			Coinpaprika: &config.DefaultExchangeRate{
				ServiceName: "coinpaprika",
				URLMask:     "https://qwe.wer",
				InRate:      "ua",
				OutRate:     "bit",
			},
		}, want: nil},
		{name: "Invalid config with empty URLMask", cfg: &config.ExchangeRate{
			Binance: &config.DefaultExchangeRate{
				ServiceName: "bitcoin",
				URLMask:     "",
				InRate:      "er",
				OutRate:     "bit",
			},
			Coinpaprika: &config.DefaultExchangeRate{
				ServiceName: "bitcoin",
				URLMask:     "https://qwe.wer",
				InRate:      "ua",
				OutRate:     "bit",
			},
		}, want: config.ErrEmptyConfig},
		{name: "Invalid config with empty InRate", cfg: &config.ExchangeRate{
			Binance: &config.DefaultExchangeRate{
				ServiceName: "bitcoin",
				URLMask:     "https://qwe.wer",
				InRate:      "",
				OutRate:     "bit",
			},
			Coinpaprika: &config.DefaultExchangeRate{
				ServiceName: "bitcoin",
				URLMask:     "https://qwe.wer",
				InRate:      "dfg",
				OutRate:     "bit",
			},
		}, want: config.ErrEmptyConfig},
		{name: "Invalid config with empty OutRate", cfg: &config.ExchangeRate{
			Binance: &config.DefaultExchangeRate{
				ServiceName: "bitcoin",
				URLMask:     "https://qwe.wer",
				InRate:      "ua",
				OutRate:     "",
			},
			Coinpaprika: &config.DefaultExchangeRate{
				ServiceName: "bitcoin",
				URLMask:     "https://qwe.wer",
				InRate:      "ua",
				OutRate:     "df",
			},
		}, want: config.ErrEmptyConfig},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cfg.Binance.Empty(); !errors.Is(err, tt.want) {
				t.Errorf("TestExchangeRate_Empty() name = %s err = %v want = %v", tt.name, err, tt.want)
			}
		})
	}
}
