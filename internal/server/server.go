package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abadojack/rtls/config"
	"github.com/rs/cors"
)

type server struct {
	router *Router
}

func newServer() *server {
	server := &server{
		router: NewRouter(),
	}

	return server
}

func RunServer() {

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "UPDATE", "OPTIONS", "DELETE", "PATCH"},
	})

	server := newServer()
	server.router.InitializeRoutes(&config.AppConfig)

	handler := c.Handler(*server.router)

	if err := http.ListenAndServe(
		fmt.Sprintf("%v:%v", "", config.AppConfig.Port),
		handler,
	); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
