package config

import (
	"os"

	"github.com/jessevdk/go-flags"
)

// Config for Assets API
type Config struct {
	Debug    bool `json:"debug" long:"debug" env:"DEBUG" description:"debug enabled"`
	DataBase DB   `json:"db"`
}

type DB struct {
	User     string
	Password string
	Host     string
	Port     int
	Name     string
}

func Init() (*Config, error) {
	var err error
	conf := &Config{}
	if err = parseFlags(conf); err != nil {
		return nil, err
	}
	return conf, nil
}

func parseFlags(opts interface{}) (err error) {
	if _, err = flags.NewParser(opts, flags.IgnoreUnknown).Parse(); err == nil {
		return
	}

	if flagsErr, ok := err.(*flags.Error); ok {
		if flagsErr.Type == flags.ErrHelp ||
			flagsErr.Type == flags.ErrRequired {
			os.Exit(0)
		}
	}

	return err
}
