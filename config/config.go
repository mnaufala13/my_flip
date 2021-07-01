package config

import (
	"github.com/caarlos0/env/v6"
	"log"
)

type Config struct {
	DBHost     string `env:"DBHost,notEmpty"`
	DBPort     string `env:"DBPort" envDefault:"5432"`
	DBUser     string `env:"DBUser,notEmpty"`
	DBPassword string `env:"DBPassword,notEmpty"`
	DBName     string `env:"DBName" envDefault:"my_market"`
	DBSsl      string `env:"DBSsl" envDefault:"disable"`
	Port       string `env:"PORT" envDefault:"8000"`
	FlipHost   string `env:"FlipHost" envDefault:"https://nextar.flip.id"`
	FlipSecret string `env:"FlipSecret,notEmpty"`
}

func NewConfig() *Config {
	cfg := &Config{}
	// Load env vars.
	if err := env.Parse(cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}
