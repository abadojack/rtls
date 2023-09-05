package main

import (
	"github.com/abadojack/rtls/config"
	"github.com/abadojack/rtls/internal/server"
)

func main() {
	// load config from .env
	config.LoadEnv()

	server.RunServer()
}
