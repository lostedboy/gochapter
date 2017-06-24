package config

import (
	"github.com/jessevdk/go-flags"
)

type Config struct {
	GoogleKey  string `long:"key" description:"Google Maps API Key" required:"true"`
}

func Parse() (*Config, error) {
	var c Config

	_, err := flags.NewParser(&c, flags.HelpFlag|flags.PassDoubleDash).Parse();

	if err != nil {
		return nil, err
	}

	return &c, nil
}