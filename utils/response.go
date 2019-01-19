package utils

import (
	"encoding/json"
	"net/http"

	"github.com/keidrun/boilerplate-gorilla-mux-for-rest-api-with-jwt/models"
)

func respond(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func Failure(w http.ResponseWriter, status int, error models.Error) {
	w.WriteHeader(status)
	respond(w, error)
}

func Success(w http.ResponseWriter, data interface{}) {
	respond(w, data)
}
