package config

import (
	"encoding/json"
	"os"

	"github.com/caarlos0/env/v10"
)

type Config struct {
	HttpServerPort string          `env:"HTTP_SERVER_PORT" envDefault:"3000"`
	Services       []ServiceConfig `json:"services"`
}

type ServiceConfig struct {
	Name        string `json:"name"`
	ContextPath string `json:"contextPath"`
	TargetURL   string `json:"targetUrl"`
}

func LoadConfig(path string) (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
