package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"

	"github.com/keidrun/boilerplate-gorilla-mux-for-rest-api-with-jwt/models"
	userRepository "github.com/keidrun/boilerplate-gorilla-mux-for-rest-api-with-jwt/repository/user"
	"github.com/keidrun/boilerplate-gorilla-mux-for-rest-api-with-jwt/utils"
	userValidator "github.com/keidrun/boilerplate-gorilla-mux-for-rest-api-with-jwt/validators/user"
)

func (c Controller) GetUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		users := make([]models.User, 0)
		var errorObj models.Error

		userRepo := userRepository.UserRepository{}
		users, err := userRepo.GetUsers(db, user, users)
		if err != nil {
			log.Println(err)
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
		id := params["id"]

		userRepo := userRepository.UserRepository{}
		user, err := userRepo.GetUser(db, user, id)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Println(err)
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

		validator := userValidator.Validator{}
		errs := validator.ValidateAddUserRequest(user)
		if len(errs) > 0 {
			var serrs []string
			for _, v := range errs {
				serrs = append(serrs, fmt.Sprintf("%v", v))
			}
			errorObj.Message = strings.Join(serrs, ",")
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
			log.Println(err)
			errorObj.Message = "Server error"
			utils.Failure(w, http.StatusInternalServerError, errorObj)
			return
		}

		utils.SuccessWithStatus(w, http.StatusCreated, user)
	}
}

func (c Controller) UpdateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		var errorObj models.Error

		params := mux.Vars(r)
		id := params["id"]

		json.NewDecoder(r.Body).Decode(&user)

		validator := userValidator.Validator{}
		errs := validator.ValidateUpdateUserRequest(user)
		if len(errs) > 0 {
			var serrs []string
			for _, v := range errs {
				serrs = append(serrs, fmt.Sprintf("%v", v))
			}
			errorObj.Message = strings.Join(serrs, ",")
			utils.Failure(w, http.StatusBadRequest, errorObj)
			return
		}

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
		userFound, err := userRepo.GetUser(db, user, id)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Println(err)
				errorObj.Message = "The user does not exist"
				utils.Failure(w, http.StatusBadRequest, errorObj)
				return
			}
			log.Println(err)
			errorObj.Message = "Server error"
			utils.Failure(w, http.StatusInternalServerError, errorObj)
			return
		}
		if user.Email == "" {
			user.Email = userFound.Email
		}
		if user.Password == "" {
			user.Password = userFound.Password
		}
		if user.Name == "" {
			user.Name = userFound.Name
		}
		if user.Age.Valid == false {
			user.Age = userFound.Age
		}
		if user.Gender.Valid == false {
			user.Gender = userFound.Gender
		}

		rowsUpdated, err := userRepo.UpdateUser(db, user, id)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Println(err)
				errorObj.Message = "The user does not exist"
				utils.Failure(w, http.StatusBadRequest, errorObj)
				return
			}
			log.Println(err)
			errorObj.Message = "Server error"
			utils.Failure(w, http.StatusInternalServerError, errorObj)
			return
		}
		if rowsUpdated != 1 {
			log.Println(err)
			errorObj.Message = "Server error"
			utils.Failure(w, http.StatusInternalServerError, errorObj)
			return
		}

		utils.Success(w, user)
	}
}

func (c Controller) RemoveUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var errorObj models.Error

		params := mux.Vars(r)
		id := params["id"]

		userRepo := userRepository.UserRepository{}
		rowsDeleted, err := userRepo.RemoveUser(db, id)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Println(err)
				errorObj.Message = "The user does not exist"
				utils.Failure(w, http.StatusBadRequest, errorObj)
				return
			}
			log.Println(err)
			errorObj.Message = "Server error"
			utils.Failure(w, http.StatusInternalServerError, errorObj)
			return
		}
		if rowsDeleted != 1 {
			log.Println(err)
			errorObj.Message = "Server error"
			utils.Failure(w, http.StatusInternalServerError, errorObj)
			return
		}

		utils.SuccessWithStatus(w, http.StatusNoContent, "")
	}
}
