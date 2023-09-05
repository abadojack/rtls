package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

const (
	secondsInOneYear = "31536000"
	ContentTypeJSON  = "application/json"
)

type Response struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}

// writeHTTPResponse would write the content type, default headers, status code, and body to the response.
// Returns error if failed. This function does not write an HTTP error so that there are no surprises
func writeHTTPResponse(w http.ResponseWriter, response interface{}, contentType string, statusCode int) error {
	setResponseHeader(w, contentType)
	w.WriteHeader(statusCode)

	// Write response
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return err
	}

	return nil
}

// setResponseHeader sets common header and content type
func setResponseHeader(w http.ResponseWriter, contentType string) {
	w.Header().Set("Content-Type", ContentTypeJSON)
	w.Header().Set("Content-Security-Policy", "default-src 'self';")
	w.Header().Set("Strict-Transport-Security", "max-age="+secondsInOneYear)
}

func WriteErrorResponse(w http.ResponseWriter, response Response, contentType string, statusCode int) {
	err := writeHTTPResponse(w, response, contentType, statusCode)
	if err != nil {
		logrus.WithError(err).Error("error sending http response")
	}
}

func unmarshallJSONFromRequest(w http.ResponseWriter, r *http.Request, data interface{}) {
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		logrus.WithError(err).Logger.Error("error decoding json")

		WriteErrorResponse(w, Response{
			Message: "error decoding json",
			Error:   err,
		}, ContentTypeJSON, http.StatusBadRequest)
		return
	}
}
