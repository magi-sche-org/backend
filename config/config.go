package config

import (
	"log"
	"net/http"

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
	CORS                 CORS   `envPrefix:"CORS_"`
	CSRF                 CSRF   `envPrefix:"CORS_"`
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

type CORS struct {
	Origins []string `env:"ORIGINS" envSeparator:"," envDefault:"http://localhost:3000"`
}
type SameSite http.SameSite
type CSRF struct {
	Disabled bool     `env:"DISABLED" envDefault:"false"`
	Domain   string   `env:"DOMAIN" envDefault:"localhost"`
	HttpOnly bool     `env:"HTTP_ONLY" envDefault:"false"`
	SameSite SameSite `env:"SAME_SITE" envDefault:"Lax"`
}

const (
	SameSiteDefaultMode SameSite = SameSite(http.SameSiteDefaultMode)
	SameSiteLaxMode     SameSite = SameSite(http.SameSiteLaxMode)
	SameSiteStrictMode  SameSite = SameSite(http.SameSiteStrictMode)
	SameSiteNoneMode    SameSite = SameSite(http.SameSiteNoneMode)
)

func (s *SameSite) UnmarshalText(text []byte) error {
	switch string(text) {
	// case "Default":
	// 	*s = SameSiteDefaultMode
	case "Lax":
		*s = SameSiteLaxMode
	case "Strict":
		*s = SameSiteStrictMode
	case "None":
		*s = SameSiteNoneMode
	default:
		// panic("invalid SameSite")
		*s = SameSiteDefaultMode
	}
	return nil
}

func New() (*Config, error) {
	config := &Config{}
	if err := env.Parse(config); err != nil {
		return nil, err
	}
	log.Printf("config: %+v", config)

	return config, nil
}
