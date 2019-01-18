package models

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name,omitempty"`
	Age      int    `json:"age,omitempty"`
	Gender   string `json:"gender,omitempty"`
}
