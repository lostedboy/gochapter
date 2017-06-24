package config

import (
	"github.com/jessevdk/go-flags"
)

type Config struct {
	GoogleKey  string `long:"key" description:"Google Maps API Key" required:"true"`
	Arguments []string
}

func Parse() (*Config, error) {
	var c Config

	args, err := flags.NewParser(&c, flags.HelpFlag|flags.PassDoubleDash).Parse();

	if err != nil {
		return nil, err
	}

	c.Arguments = args

	return &c, nil
}