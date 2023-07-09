package model_test

import (
	"testing"

	"bitcoin_checker_api/internal/model"
)

func TestNewExchangeRate(t *testing.T) {
	type want struct {
		base  string
		quote string
		price float64
	}
	tests := []struct {
		name  string
		base  string
		quote string
		price float64
		want  want
	}{
		{name: "Valid exchange rate model creation", base: "Bitcoin", quote: "Ukrainian Hryvnia", price: 20000.0202, want: want{base: "Bitcoin", quote: "Ukrainian Hryvnia", price: 20000.0202}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tER := model.NewExchangeRate(tt.base, tt.quote, tt.price)
			if tER.BaseCurrency != tt.base || tER.QuoteCurrency != tt.quote || tER.Price != tt.price {
				t.Errorf("TestNewUser() name = %v got = %s want = %s", tt.name, tER.BaseCurrency, tt.base)
			}
		})
	}
}
