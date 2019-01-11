package userRepository

import (
	"database/sql"
	"log"

	"github.com/keidrun/boilerplate-golang-for-rest-api/models"
)

type UserRepository struct{}

func (u UserRepository) Signup(db *sql.DB, user models.User) (models.User, error) {
	err := db.QueryRow("insert into users (email, password) values($1, $2) RETURNING id;", user.Email, user.Password).Scan(&user.ID)

	if err != nil {
		log.Println(err)
		return user, err
	}

	user.Password = ""
	return user, nil
}

func (u UserRepository) Login(db *sql.DB, user models.User) (models.User, error) {
	row := db.QueryRow("select * from users where email=$1", user.Email)
	err := row.Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		log.Println(err)
		return user, err
	}

	return user, nil
}
