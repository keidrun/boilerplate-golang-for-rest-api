package utils

import (
	"log"
	"os"

	"github.com/keidrun/boilerplate-gorilla-mux-for-rest-api-with-jwt/models"

	jwt "github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"
)

func ComparePasswords(hashedPassword string, password []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), password)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func GenerateToken(user models.User) (string, error) {
	var err error
	secret := os.Getenv("SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   os.Getenv("ISSUER"),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Println(err)
		return tokenString, err
	}

	return tokenString, nil
}
