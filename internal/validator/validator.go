package validator

import (
	"errors"
	"net/mail"
	"net/url"
)

var (
	ErrInvalidMailAddress = errors.New("invalid mail address")
	ErrInvalidURL         = errors.New("invalid url")
)

func ValidMailAddress(address string) error {
	if _, err := mail.ParseAddress(address); err != nil {
		return ErrInvalidMailAddress
	}
	return nil
}

func ValidURL(rawURL string) error {
	if _, err := url.ParseRequestURI(rawURL); err != nil {
		return ErrInvalidURL
	}
	return nil
}
