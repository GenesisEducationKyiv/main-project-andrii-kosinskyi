package mapper

import "bitcoin_checker_api/internal/model"

type Mapper interface {
	Map([]byte) (*model.ExchangeRate, error)
	Name() string
}
