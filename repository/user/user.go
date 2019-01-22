package userRepository

import (
	"database/sql"

	"github.com/keidrun/boilerplate-gorilla-mux-for-rest-api-with-jwt/models"
)

func (u UserRepository) GetUsers(db *sql.DB, user models.User, users []models.User) ([]models.User, error) {
	rows, err := db.Query("select id,email,password,name,age,gender from users")
	if err != nil {
		return users, err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.Age, &user.Gender)
		if err != nil {
			return users, err
		}

		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return users, err
	}

	return users, nil
}

func (u UserRepository) GetUser(db *sql.DB, user models.User, id string) (models.User, error) {
	rows := db.QueryRow("select id,email,password,name,age,gender from users where id=$1", id)
	err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.Age, &user.Gender)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u UserRepository) AddUser(db *sql.DB, user models.User) (models.User, error) {
	err := db.QueryRow(
		"insert into users (email, password, name, age, gender) values($1, $2, $3, $4, $5) RETURNING id;",
		user.Email, user.Password, user.Name, user.Age, user.Gender).Scan(&user.ID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u UserRepository) UpdateUser(db *sql.DB, user models.User, id string) (int64, error) {
	result, err := db.Exec(
		"update users set email=$1, password=$2, name=$3, age=$4, gender=$5 where id=$6 RETURNING id",
		&user.Email, &user.Password, &user.Name, &user.Age, &user.Gender, id)
	if err != nil {
		return 0, err
	}

	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsUpdated, nil
}

func (u UserRepository) RemoveUser(db *sql.DB, id string) (int64, error) {
	result, err := db.Exec("delete from users where id = $1", id)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsDeleted, nil
}
