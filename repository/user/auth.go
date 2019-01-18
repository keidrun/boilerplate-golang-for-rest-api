package userRepository

import (
	"database/sql"
	"log"

	"github.com/keidrun/boilerplate-gorilla-mux-for-rest-api-with-jwt/models"
)

func (u UserRepository) Signup(db *sql.DB, user models.User) (models.User, error) {
	err := db.QueryRow("insert into users (email, password, name, age, gender) values($1, $2, $3, $4, $5) RETURNING id;", user.Email, user.Password, user.Name, user.Age, user.Gender).Scan(&user.ID)

	if err != nil {
		log.Println(err)
		return user, err
	}

	user.Password = ""
	return user, nil
}

func (u UserRepository) Login(db *sql.DB, user models.User) (models.User, error) {

	row := db.QueryRow("select id,email,password,name,age,gender from users where email=$1", user.Email)
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.Age, &user.Gender)

	if err != nil {
		log.Println(err)
		return user, err
	}

	return user, nil
}
