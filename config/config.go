package config

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	// Environment running i.e. dev, stage, prod
	Environment string `env:"ENVIRONMENT" envDefault:"dev"`

	// Port for http app
	Port uint32 `env:"PORT" envDefault:"3000"`

	// DB config format: "user:pass@host:port/dbname?charset=utf8"
	DB string `env:"DB,required"`
}

var AppConfig Config

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file. Defaulting to environment")
	}

	err = env.Parse(&AppConfig)
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}
}
