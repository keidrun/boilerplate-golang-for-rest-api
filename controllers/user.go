package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"

	"github.com/keidrun/boilerplate-gorilla-mux-for-rest-api-with-jwt/models"
	userRepository "github.com/keidrun/boilerplate-gorilla-mux-for-rest-api-with-jwt/repository/user"
	"github.com/keidrun/boilerplate-gorilla-mux-for-rest-api-with-jwt/utils"
)

func (c Controller) GetUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		var users []models.User
		var errorObj models.Error

		userRepo := userRepository.UserRepository{}

		users, err := userRepo.GetUsers(db, user, users)
		if err != nil {
			errorObj.Message = "Server error"
			utils.Failure(w, http.StatusInternalServerError, errorObj)
			return
		}

		utils.Success(w, users)
	}
}

func (c Controller) GetUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		var errorObj models.Error

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			errorObj.Message = "\"id\" is wrong"
			utils.Failure(w, http.StatusBadRequest, errorObj)
			return
		}

		userRepo := userRepository.UserRepository{}
		user, err = userRepo.GetUser(db, user, id)
		if err != nil {
			if err == sql.ErrNoRows {
				errorObj.Message = "The user does not exist"
				utils.Failure(w, http.StatusBadRequest, errorObj)
				return
			}
			log.Println(err)
			errorObj.Message = "Server error"
			utils.Failure(w, http.StatusInternalServerError, errorObj)
			return
		}

		utils.Success(w, user)
	}
}

func (c Controller) AddUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		var errorObj models.Error

		json.NewDecoder(r.Body).Decode(&user)
		if user.Email == "" {
			errorObj.Message = "\"email\" is missing"
			utils.Failure(w, http.StatusBadRequest, errorObj)
			return
		}
		if user.Password == "" {
			errorObj.Message = "\"password\" is missing"
			utils.Failure(w, http.StatusBadRequest, errorObj)
			return
		}
		if user.Name == "" {
			errorObj.Message = "\"name\" is missing."
			utils.Failure(w, http.StatusBadRequest, errorObj)
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			errorObj.Message = "Something went wrong when generating token"
			utils.Failure(w, http.StatusBadRequest, errorObj)
			return
		}
		user.Password = string(hash)

		userRepo := userRepository.UserRepository{}
		user, err = userRepo.AddUser(db, user)
		if err != nil {
			errorObj.Message = "Server error"
			utils.Failure(w, http.StatusInternalServerError, errorObj)
			return
		}

		utils.Success(w, user)
	}
}

func (c Controller) UpdateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		var errorObj models.Error

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			errorObj.Message = "\"id\" is wrong"
			utils.Failure(w, http.StatusBadRequest, errorObj)
			return
		}

		json.NewDecoder(r.Body).Decode(&user)
		if user.Password != "" {
			hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
			if err != nil {
				errorObj.Message = "Something went wrong when generating token"
				utils.Failure(w, http.StatusBadRequest, errorObj)
				return
			}
			user.Password = string(hash)
		}

		userRepo := userRepository.UserRepository{}
		rowsUpdated, err := userRepo.UpdateUser(db, user, id)
		if err != nil {
			if err == sql.ErrNoRows {
				errorObj.Message = "The user does not exist"
				utils.Failure(w, http.StatusBadRequest, errorObj)
				return
			}
			log.Println(err)
			errorObj.Message = "Server error"
			utils.Failure(w, http.StatusInternalServerError, errorObj)
			return
		}

		utils.Success(w, rowsUpdated)
	}
}

func (c Controller) RemoveUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var errorObj models.Error

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			errorObj.Message = "\"id\" is wrong"
			utils.Failure(w, http.StatusBadRequest, errorObj)
			return
		}

		userRepo := userRepository.UserRepository{}
		rowsDeleted, err := userRepo.RemoveUser(db, id)
		if err != nil {
			if err == sql.ErrNoRows {
				errorObj.Message = "The user does not exist"
				utils.Failure(w, http.StatusBadRequest, errorObj)
				return
			}
			log.Println(err)
			errorObj.Message = "Server error"
			utils.Failure(w, http.StatusInternalServerError, errorObj)
			return
		}

		utils.Success(w, rowsDeleted)
	}
}
