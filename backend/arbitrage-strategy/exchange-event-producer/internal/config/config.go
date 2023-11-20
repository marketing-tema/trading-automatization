package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServerPort string `envconfig:"SERVER_PORT" default:"8080"`
	MqHost string `envconfig:"MQ_HOST" required:"true"`
	MqPort string `envconfig:"MQ_PORT" required:"true"`
}

func New() (*Config, error) {
	var config Config
	err := envconfig.Process("myapp", &config)
	return &config, err
}
