package mapper

import (
	"encoding/json"
	"time"

	"bitcoin_checker_api/internal/model"
)

type symbol struct {
	name          string
	baseFullName  string
	quoteFullName string
}

type BinanceMapper struct {
	name    string
	symbols map[string]*symbol
}

type BinanceRespBody struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price,string"`
}

func NewBinanceMapper() Mapper {
	symbols := make(map[string]*symbol)
	symbols["BTCUAH"] = &symbol{
		name:          "BTCUAH",
		baseFullName:  "Bitcoin",
		quoteFullName: "Ukrainian Hryvnia",
	}
	return &BinanceMapper{
		name:    BinanceService,
		symbols: symbols,
	}
}

func (that *BinanceMapper) Name() string {
	return that.name
}

func (that *BinanceMapper) Map(serviceRespBody []byte) (*model.ExchangeRate, error) {
	brb := &BinanceRespBody{}
	if err := json.Unmarshal(serviceRespBody, brb); err != nil {
		return nil, ErrUnmarshal
	}
	symbolResolver, err := that.resolveSymbol(brb.Symbol)
	if err != nil {
		return nil, err
	}
	return &model.ExchangeRate{
		BaseCurrency:  symbolResolver.baseFullName,
		QuoteCurrency: symbolResolver.quoteFullName,
		Price:         brb.Price,
		Date:          time.Now(),
	}, nil
}

func (that *BinanceMapper) resolveSymbol(symbol string) (*symbol, error) {
	s, ok := that.symbols[symbol]
	if !ok {
		return nil, ErrBinanceSymbol
	}
	return s, nil
}
