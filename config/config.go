package config

import (
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Service struct {
	Port int64 `toml:"port"`
}

type Converter struct {
	Endpoint string `toml:"endpoint"`
}

type InternalStorage struct {
	Path string `toml:"path"`
}

type Config struct {
	Service         Service
	Converter       Converter
	InternalStorage InternalStorage
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

	return nil
}

func NewConfig() *Config {
	return &Config{
		Service:         Service{},
		Converter:       Converter{},
		InternalStorage: InternalStorage{},
	}
}
