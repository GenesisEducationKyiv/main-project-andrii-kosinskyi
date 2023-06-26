package config

import (
	"fmt"
	"log"
	"os"

	"bitcoin_checker_api/internal/validator"

	"github.com/pelletier/go-toml/v2"
)

type Server struct {
	Port int64 `toml:"port"`
}

type ExchangeRate struct {
	URLMask string `toml:"url_mask"`
	InRate  string `toml:"in_rate_name"`
	OutRate string `toml:"out_rate_name"`
}

type Storage struct {
	Path string `toml:"path"`
}

type EmailService struct {
	APIKey      string `toml:"api_key"`
	FromAddress string `toml:"from_address"`
	FromName    string `toml:"from_name"`
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
	rawURL := fmt.Sprintf(that.ExchangeRate.URLMask, that.ExchangeRate.InRate, that.ExchangeRate.OutRate)
	if _, err := validator.ValidURLWithError(rawURL); err != nil {
		return err
	}
	return nil
}
