package config

import "github.com/caarlos0/env/v8"

type Config struct {
	App App
	DB  DB
	Log Log
}

type App struct {
	Name      string `env:"APP_NAME"`
	Port      string `env:"APP_PORT" envDefault:":8080"`
	JWTSecret string `env:"JWT_SECRET"`
}

type DB struct {
	Host    string `env:"DB_HOST" envDefault:"localhost"`
	Port    string `env:"DB_PORT" envDefault:"5432"`
	User    string `env:"DB_USER"`
	Pwd     string `env:"DB_PWD"`
	Name    string `env:"DB_NAME"`
	SSLMode string `env:"DB_SSL_MODE" envDefault:"disable"`
	DNS     string `env:"DB_DNS"`
}

type Log struct {
	Level string `env:"LOG_LEVEL" envDefault:"debug"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
