package exchange_rate

import "context"

type ExchangeRater interface {
	Get(ctx context.Context) (string, error)
}
