package exchangerate

import "context"

type ExchangeRater interface {
	Get(ctx context.Context) ([]byte, error)
}
