package mapper

const CoinPaprikaService = "coinpaprika"

func NewExchangeRateMapper(serviceName string) (Mapper, error) {
	switch serviceName {
	case CoinPaprikaService:
		return NewCoinPaprikaMapper(), nil
	default:
		return nil, ErrUnknownService
	}
}
