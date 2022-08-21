package config

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Logging     LogConfig         `envconfig:"LOGGING"`
	Rest        RestConfig        `envconfig:"REST"`
	Persistence PersistenceConfig `envconfig:"PERSISTENCE"`
}

func NewConfig() *Config {
	var config Config
	err := envconfig.Process("RUNNING", &config)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &config
}

type LogConfig struct {
	Level string `default:"info"`
}

type RestConfig struct {
	Port                  string `envconfig:"PORT"`
	UsePasswordMiddleware bool   `split_words:"true" default:"true"`

	UseCorsMiddleware bool `split_words:"true" default:"true"`
}

type PersistenceConfig struct {
	StorerType string `split_words:"true" default:"mysql"`
	Address    string `envconfig:"DATABASE_URL"`
}
