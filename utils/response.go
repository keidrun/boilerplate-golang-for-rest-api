package utils

import (
	"encoding/json"
	"net/http"

	"github.com/keidrun/boilerplate-gorilla-mux-for-rest-api-with-jwt/models"
)

func respond(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func Failure(w http.ResponseWriter, status int, error models.Error) {
	respond(w, status, error)
}

func Success(w http.ResponseWriter, data interface{}) {
	respond(w, http.StatusOK, data)
}

func SuccessWithStatus(w http.ResponseWriter, status int, data interface{}) {
	respond(w, status, data)
}
