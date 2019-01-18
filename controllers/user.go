package controllers

import (
	"net/http"

	"github.com/keidrun/boilerplate-gorilla-mux-for-rest-api-with-jwt/utils"
)

func (c Controller) GetUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.Respond(w, "GET Users")
	}
}
