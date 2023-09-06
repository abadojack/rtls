package main

import (
	"time"

	"github.com/abadojack/rtls/config"
	"github.com/abadojack/rtls/internal/db"
	"github.com/abadojack/rtls/internal/models"
	"github.com/abadojack/rtls/internal/server"
	"github.com/avast/retry-go"
	"github.com/sirupsen/logrus"
)

func main() {
	// load config from .env
	config.LoadEnv()

	err := retry.Do(
		func() error {
			return Migrate()
		},
		retry.Delay(2*time.Second),
		retry.Attempts(5),
	)
	if err != nil {
		logrus.WithError(err).Fatal("could not run db migrations")
		return
	}

	server.RunServer()
}

func Migrate() error {
	db, err := db.GetDB()
	if err != nil {
		return err
	}

	logrus.Infoln(
		"running db migrations",
	)

	return db.AutoMigrate(&models.Player{})
}
