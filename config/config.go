package config

import (
	"fmt"
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

func New() *Config {
	cfg := &Config{}

	configPath := "config.yml"

	if configPath == "" {
		fmt.Errorf("Config path is not set")
		return nil
	}

	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		fmt.Errorf("Cannot read config: %s", err)
		return nil
	}

	return cfg

}
