package config

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServerPort int    `envconfig:"SERVER_PORT" default:"8080"`
	MqHost     string `envconfig:"MQ_HOST" required:"true"`
	MqPort     string `envconfig:"MQ_PORT" required:"true"`
}

func New() (*Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err := config.validate(); err != nil {
		return nil, err
	}
	return &config, err
}

func (c Config) validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.MqHost, validation.Required),
		validation.Field(&c.MqPort, validation.Required),
	)
}
