package exchangerate

import (
	"context"

	"bitcoin_checker_api/internal/model"
)

type ExchangeRater interface {
	Get(ctx context.Context) (*model.ExchangeRate, error)
	SetNext(next ExchangeRater)
}
