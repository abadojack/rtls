package server

import (
	"net/http"

	"github.com/abadojack/rtls/config"
	"github.com/abadojack/rtls/internal/server/handlers"
	"github.com/gorilla/mux"
)

// Router wraps gorilla mux router so that we're able to create methods from it
type Router struct {
	*mux.Router
}

func NewRouter() *Router {
	return &Router{mux.NewRouter()}
}

// InitializeRoutes Initializa routes and register handlers
func (r *Router) InitializeRoutes(cfg *config.Config) {
	r.HandleFunc("/healthcheck", handlers.HealthCheckHandler).
		Methods(http.MethodGet).
		Name("healthcheck")

	r.HandleFunc("/players", handlers.PostPlayersHandler).
		Methods(http.MethodPost).
		Name("postPlayers")

	r.HandleFunc("/players/{id}/rank", handlers.GetPlayersRankHandler).
		Methods(http.MethodGet).
		Name("getPlayersRank")

	r.HandleFunc("/players/{id}/score", handlers.PutPlayersHandler).
		Methods(http.MethodPut).
		Name("putPlayersScore")

	r.HandleFunc("/leaderboard/top/{N}", handlers.GetLeaderBoardHandler).
		Methods(http.MethodGet).
		Name("leaderboardTopN")
}
