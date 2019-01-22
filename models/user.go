package models

import (
	"encoding/json"
)

type User struct {
	ID       string     `json:"id"`
	Email    string     `json:"email"`
	Password string     `json:"password,omitempty"`
	Name     string     `json:"name,omitempty"`
	Age      NullInt64  `json:"age,omitempty"`
	Gender   NullString `json:"gender,omitempty"`
}

func (u User) MarshalJSON() ([]byte, error) {
	var response struct {
		ID     string     `json:"id"`
		Email  string     `json:"email"`
		Name   string     `json:"name,omitempty"`
		Age    NullInt64  `json:"age,omitempty"`
		Gender NullString `json:"gender,omitempty"`
	}
	response.ID = u.ID
	response.Email = u.Email
	response.Name = u.Name
	response.Age = u.Age
	response.Gender = u.Gender

	return json.Marshal(&response)
}
