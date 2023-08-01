package mapper

import "errors"

var (
	ErrUnknownService = errors.New("unknown exchange rate service name")
	ErrUnmarshal      = errors.New("unknown api response")
)
