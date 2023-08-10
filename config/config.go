package config

import (
	"log"

	"github.com/caarlos0/env/v9"
)

type Config struct {
	Port                 int    `env:"PORT" envDefault:"80"`
	Env                  string `env:"ENV" envDefault:"dev"`
	SecretKey            string `env:"SECRET_KEY" envDefault:"secret"`
	AccessExpireMinutes  int    `env:"ACCESS_TOKEN_EXPIRE_MINUTES" envDefault:"5"`
	RefreshExpireMinutes int    `env:"REFRESH_TOKEN_EXPIRE_MINUTES" envDefault:"43200"`
	MySQL                MySQL  `envPrefix:"MYSQL_"`
	OAuth                OAuth  `envPrefix:"OAUTH_"`
}

type MySQL struct {
	Host     string `env:"HOST"`
	Port     int    `env:"PORT"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	DBName   string `env:"DATABASE"`
}
type OAuth struct {
	Google Client `envPrefix:"GOOGLE_"`
}
type Client struct {
	ClientID     string `env:"CLIENT_ID"`
	ClientSecret string `env:"CLIENT_SECRET"`
}

func New() (*Config, error) {
	config := &Config{}
	if err := env.Parse(config); err != nil {
		return nil, err
	}
	log.Printf("config: %+v", config)

	return config, nil
}
