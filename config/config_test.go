package config

import "testing"

func TestEmailService_Empty(t *testing.T) {

	tests := []struct {
		name string
		cfg  *EmailService
		want error
	}{
		{name: "Valid config", cfg: &EmailService{
			APIKey:      "qwer-3423d-sdfasdf",
			FromAddress: "tarsa@schecvcnko.com",
			FromName:    "tarsa@schecvcnko.com",
		}, want: nil},
		{name: "Invalid config with empty APIKey", cfg: &EmailService{
			APIKey:      "",
			FromAddress: "tarsa@schecvcnko.com",
			FromName:    "tarsa@schecvcnko.com",
		}, want: ErrEmptyConfig},
		{name: "Invalid config with empty FromAddress", cfg: &EmailService{
			APIKey:      "qwer-3423d-sdfasdf",
			FromAddress: "",
			FromName:    "tarsa@schecvcnko.com",
		}, want: ErrEmptyConfig},
		{name: "Invalid config with empty FromName", cfg: &EmailService{
			APIKey:      "qwer-3423d-sdfasdf",
			FromAddress: "tarsa@schecvcnko.com",
			FromName:    "",
		}, want: ErrEmptyConfig},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cfg.Empty(); err != tt.want {
				t.Errorf("TestEmailService_Empty() name = %s err = %v want = %v", tt.name, err, tt.want)
			}
		})
	}

}

func TestExchangeRate_Empty(t *testing.T) {
	tests := []struct {
		name string
		cfg  *ExchangeRate
		want error
	}{
		{name: "Valid config", cfg: &ExchangeRate{
			URLMask: "https://qwe.wer",
			InRate:  "ua",
			OutRate: "bit",
		}, want: nil},
		{name: "Invalid config with empty URLMask", cfg: &ExchangeRate{
			URLMask: "",
			InRate:  "ua",
			OutRate: "bit",
		}, want: ErrEmptyConfig},
		{name: "Invalid config with empty InRate", cfg: &ExchangeRate{
			URLMask: "https://qwe.wer",
			InRate:  "",
			OutRate: "bit",
		}, want: ErrEmptyConfig},
		{name: "Invalid config with empty OutRate", cfg: &ExchangeRate{
			URLMask: "https://qwe.wer",
			InRate:  "ua",
			OutRate: "",
		}, want: ErrEmptyConfig},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cfg.Empty(); err != tt.want {
				t.Errorf("TestExchangeRate_Empty() name = %s err = %v want = %v", tt.name, err, tt.want)
			}
		})
	}
}
