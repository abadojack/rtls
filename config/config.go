package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	// Environment running i.e. dev, stage, prod
	Environment string `env:"ENVIRONMENT" envDefault:"dev"`

	// Port for http app
	Port uint32 `env:"PORT" envDefault:"8085"`

	// DB config format: "user:pass@host:port/dbname?charset=utf8"
	DBName     string `env:"DB_NAME,required"`
	DBUser     string `env:"DB_USER,required"`
	DBHost     string `env:"DB_HOST,required"`
	DBPassword string `env:"DB_PASSWORD,required"`

	// Redis config
	RedisHost     string `env:"REDIS_HOST" envDefault:"redis:6379"`
	RedisPassword string `env:"REDIS_PASSWORD"`
}

var AppConfig Config

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		logrus.WithError(err).Println("Error loading .env file. Defaulting to environment")
	}

	err = env.Parse(&AppConfig)
	if err != nil {
		logrus.WithError(err).Fatalf("unable to parse ennvironment variables: %e", err)
	}
}
