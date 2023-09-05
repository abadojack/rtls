package handlers

import (
	"net/http"

	"github.com/abadojack/rtls/internal/models"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type PlayerRequest struct {
	Score int `json:"score"`
}

type PlayerResponse struct {
	ID    string `json:"id"`
	Score int    `json:"score"`
}

// Validate request data
func (p *PlayerRequest) Validate() error {
	// nothing to validate currently
	return nil
}

func PostPlayersHandler(w http.ResponseWriter, r *http.Request) {
	var request PlayerRequest
	unmarshallJSONFromRequest(w, r, &request)

	err := request.Validate()
	if err != nil {
		logrus.WithError(err).Error("request validation error")
		WriteErrorResponse(w, Response{Message: "request validation error", Error: err}, ContentTypeJSON, http.StatusInternalServerError)
		return
	}

	player, err := models.NewPlayer(request.Score)
	if err != nil {
		logrus.WithError(err).Error("could not write player to the database")
		WriteErrorResponse(w, Response{Message: "could not write player to the database", Error: err}, ContentTypeJSON, http.StatusInternalServerError)
		return
	}

	p := PlayerResponse{
		ID:    player.ID,
		Score: player.Score,
	}

	err = writeHTTPResponse(w, p, ContentTypeJSON, http.StatusCreated)
	if err != nil {
		logrus.WithError(err).Error("error writing response")
		WriteErrorResponse(w, Response{Message: "error writing response", Error: err}, ContentTypeJSON, http.StatusInternalServerError)
	}
}

func PutPlayersHandler(w http.ResponseWriter, r *http.Request) {
	var request PlayerRequest
	unmarshallJSONFromRequest(w, r, &request)

	err := request.Validate()
	if err != nil {
		logrus.WithError(err).Error("request validation error")
		WriteErrorResponse(w, Response{Message: "request validation error", Error: err}, ContentTypeJSON, http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)

	playerID := vars["id"]

	player := models.Player{
		ID:    playerID,
		Score: request.Score,
	}

	err = player.Update()
	if err != nil {
		logrus.WithError(err).Error("could not update player in the database")
		WriteErrorResponse(w, Response{Message: "could not update player in the database", Error: err}, ContentTypeJSON, http.StatusInternalServerError)
		return
	}

	p := PlayerResponse{
		ID:    player.ID,
		Score: player.Score,
	}

	err = writeHTTPResponse(w, p, ContentTypeJSON, http.StatusOK)
	if err != nil {
		logrus.WithError(err).Error("error writing response")
		WriteErrorResponse(w, Response{Message: "error writing response", Error: err}, ContentTypeJSON, http.StatusInternalServerError)
	}
}

type PlayerRankResponse struct {
	ID    string `json:"id"`
	Score int    `json:"score"`
	Rank  int    `json:"rank"`
}

func GetPlayersRankHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	playerID := vars["id"]

	player := &models.Player{
		ID: playerID,
	}

	player, err := player.Get()
	if err != nil {
		logrus.WithError(err).Error("could not get player from the database")
		WriteErrorResponse(w, Response{Message: "could not get player from the database", Error: err}, ContentTypeJSON, http.StatusInternalServerError)
		return
	}

	p := PlayerRankResponse{
		ID:    player.ID,
		Score: player.Score,
		Rank:  player.Rank,
	}

	err = writeHTTPResponse(w, p, ContentTypeJSON, http.StatusOK)
	if err != nil {
		logrus.WithError(err).Error("error writing response")
		WriteErrorResponse(w, Response{Message: "error writing response", Error: err}, ContentTypeJSON, http.StatusInternalServerError)
	}

}
