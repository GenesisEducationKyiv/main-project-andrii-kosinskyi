package config

import (
	"fmt"
	"log"
	"os"
)
import "github.com/pelletier/go-toml/v2"

type Converter struct {
	Endpoint string `toml:"endpoint"`
}

type InternalStorage struct {
	Path string `toml:"path"`
}

type Config struct {
	Converter       Converter
	InternalStorage InternalStorage
}

func (that *Config) Load() error {
	f, err := os.ReadFile("../_env/example.toml")
	if err != nil {
		// failed to create/open the file
		log.Fatal(err)
		return err
	}
	fmt.Println("HI")
	if err := toml.Unmarshal(f, that); err != nil {
		// failed to encode
		log.Fatal(err)
		return err
	}
	fmt.Println("Bue")
	return nil
}

func NewConfig() *Config {
	return &Config{
		Converter:       Converter{},
		InternalStorage: InternalStorage{},
	}
}
