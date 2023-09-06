package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/abadojack/rtls/internal/models"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func GetLeaderBoardHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	numberOfPlayers := vars["N"]

	n, err := strconv.Atoi(numberOfPlayers)
	if err != nil {
		logrus.WithError(err).Error(fmt.Sprintf("couldn't convert path param %s from string to integer", numberOfPlayers))
		WriteErrorResponse(w, Response{Message: fmt.Sprintf("couldn't convert path param %s from string to integer", numberOfPlayers), Error: err}, ContentTypeJSON, http.StatusBadRequest)
		return
	}

	players, err := models.GetLeaderBoard(n)
	if err != nil {
		logrus.WithError(err).Error(fmt.Sprintf("couldn't get top %d players", n))
		WriteErrorResponse(w, Response{Message: fmt.Sprintf("couldn't get top %d players", n), Error: err}, ContentTypeJSON, http.StatusInternalServerError)
		return
	}

	err = writeHTTPResponse(w, players, ContentTypeJSON, http.StatusOK)
	if err != nil {
		logrus.WithError(err).Error("error writing response")
		WriteErrorResponse(w, Response{Message: "error writing response", Error: err}, ContentTypeJSON, http.StatusInternalServerError)
	}
}
