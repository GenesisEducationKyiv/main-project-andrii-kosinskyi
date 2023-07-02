package config

import (
	"errors"
	"fmt"
	"log"
	"os"

	"bitcoin_checker_api/internal/validator"

	"github.com/pelletier/go-toml/v2"
)

var ErrEmptyConfig = errors.New("one or more fields empty in config")

type EmptyConfigChecker interface {
	Empty() error
}

type Server struct {
	Port int64 `toml:"port"`
}

type ExchangeRate struct {
	URLMask string `toml:"url_mask"`
	InRate  string `toml:"in_rate_name"`
	OutRate string `toml:"out_rate_name"`
	EmptyConfigChecker
}

func (that *ExchangeRate) Empty() error {
	if that.URLMask == "" || that.InRate == "" || that.OutRate == "" {
		return ErrEmptyConfig
	}
	return nil
}

type Storage struct {
	Path string `toml:"path"`
}

type EmailService struct {
	APIKey      string `toml:"api_key"`
	FromAddress string `toml:"from_address"`
	FromName    string `toml:"from_name"`
	EmptyConfigChecker
}

func (that *EmailService) Empty() error {
	if that.APIKey == "" || that.FromName == "" || that.FromAddress == "" {
		return ErrEmptyConfig
	}
	return nil
}

func NewConfig() *Config {
	return &Config{
		Server:       Server{},
		ExchangeRate: ExchangeRate{},
		Storage:      Storage{},
		EmailService: EmailService{},
	}
}

type Config struct {
	Server       Server
	ExchangeRate ExchangeRate
	Storage      Storage
	EmailService EmailService
}

func (that *Config) Load() error {
	f, err := os.ReadFile("./_env/example.toml")
	if err != nil {
		// failed to create/open the file
		log.Fatal(err)
		return err
	}

	if err = toml.Unmarshal(f, that); err != nil {
		// failed to encode
		log.Fatal(err)
		return err
	}

	if err = that.validConfig(); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (that *Config) validConfig() error {
	if err := that.EmailService.Empty(); err != nil {
		return err
	}

	if err := that.ExchangeRate.Empty(); err != nil {
		return err
	}

	rawURL := fmt.Sprintf(that.ExchangeRate.URLMask, that.ExchangeRate.InRate, that.ExchangeRate.OutRate)
	if err := validator.ValidURL(rawURL); err != nil {
		return err
	}
	return nil
}
