package handlers

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type HealthCheckResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

// HealthCheckHandler exists to know if the application is running
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {

	healthCheckResponse := HealthCheckResponse{
		Status:  "online",
		Version: "0.1",
	}

	err := writeHTTPResponse(w, healthCheckResponse, ContentTypeJSON, http.StatusOK)
	if err != nil {
		logrus.WithError(err).Error("json encode error")
	}
}
