package config

import (
	"errors"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env-default:"dev"`
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Port        string        `yaml:"port" env:"TODO_PORT" env-default:"7540"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle-timeout" env-default:"60s"`
}

func New() (*Config, error) {
	cfg := &Config{}

	configPath := "config.yml"

	if configPath == "" {
		return nil, errors.New("Config path is not set")
	}

	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		return nil, errors.New("Cannot read config")
	}

	return cfg, nil

}
