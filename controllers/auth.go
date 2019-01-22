package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	userValidator "github.com/keidrun/boilerplate-gorilla-mux-for-rest-api-with-jwt/validators/user"

	"github.com/keidrun/boilerplate-gorilla-mux-for-rest-api-with-jwt/config"

	"github.com/keidrun/boilerplate-gorilla-mux-for-rest-api-with-jwt/models"
	userRepository "github.com/keidrun/boilerplate-gorilla-mux-for-rest-api-with-jwt/repository/user"
	"github.com/keidrun/boilerplate-gorilla-mux-for-rest-api-with-jwt/utils"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func (c Controller) Signup(db *sql.DB) http.HandlerFunc {
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
		user, err = userRepo.Signup(db, user)
		if err != nil {
			log.Println(err)
			errorObj.Message = "Server error"
			utils.Failure(w, http.StatusInternalServerError, errorObj)
			return
		}

		utils.SuccessWithStatus(w, http.StatusCreated, user)
	}
}

func (c Controller) Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		var jwt models.JWT
		var errorObj models.Error

		json.NewDecoder(r.Body).Decode(&user)

		if user.Email == "" {
			errorObj.Message = "\"email\" is missing."
			utils.Failure(w, http.StatusBadRequest, errorObj)
			return
		}

		if user.Password == "" {
			errorObj.Message = "\"password\" is missing."
			utils.Failure(w, http.StatusBadRequest, errorObj)
			return
		}

		password := user.Password

		userRepo := userRepository.UserRepository{}
		user, err := userRepo.Login(db, user)
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

		token, err := utils.GenerateToken(user)
		if err != nil {
			log.Println(err)
			errorObj.Message = "Server error"
			utils.Failure(w, http.StatusInternalServerError, errorObj)
			return
		}

		hashedPassword := user.Password
		isValidPassword := utils.ComparePasswords(hashedPassword, []byte(password))
		if isValidPassword {
			w.Header().Set("Authorization", token)
			jwt.Token = token
			utils.Success(w, jwt)
		} else {
			errorObj.Message = "Invalid password."
			utils.Failure(w, http.StatusUnauthorized, errorObj)
		}
	}
}

func (c Controller) TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conf := config.GetConfig()
		var errorObj models.Error

		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) == 2 {
			authToken := bearerToken[1]

			token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Something went wrong when parsing token")
				}
				return []byte(conf.JwtSecret), nil
			})

			if err != nil {
				errorObj.Message = err.Error()
				utils.Failure(w, http.StatusUnauthorized, errorObj)
				return
			}

			if token.Valid {
				next.ServeHTTP(w, r)
			} else {
				errorObj.Message = err.Error()
				utils.Failure(w, http.StatusUnauthorized, errorObj)
				return
			}
		} else {
			errorObj.Message = "Invalid token"
			utils.Failure(w, http.StatusUnauthorized, errorObj)
			return
		}
	})
}
