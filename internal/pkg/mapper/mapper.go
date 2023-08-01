package mapper

const (
	CoinPaprikaService = "coinpaprika"
	BinanceService     = "binance"
)

func NewExchangeRateMapper(serviceName string) (Mapper, error) {
	switch serviceName {
	case CoinPaprikaService:
		return NewCoinPaprikaMapper(), nil
	case BinanceService:
		return NewBinanceMapper(), nil
	default:
		return nil, ErrUnknownService
	}
}
