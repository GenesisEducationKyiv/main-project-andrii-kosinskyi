package exchange_rate

type ExchangeRater interface {
	Get() (string, error)
}
