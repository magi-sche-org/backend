package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Env        string `env:"SCHE_ENV" envDefault:"dev"`
	Port       int    `env:"PORT" envDefault:"50051"`
	DBHost     string `env:"MYSQL_HOST" envDefault:"127.0.0.1"`
	DBPort     int    `env:"MYSQL_PORT" envDefault:"33306"`
	DBUser     string `env:"MYSQL_USER" envDefault:"mysql"`
	DBPassword string `env:"MYSQL_PASSWORD" envDefault:"mysql"`
	DBName     string `env:"MYSQL_DATABASE" envDefault:"magische"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
