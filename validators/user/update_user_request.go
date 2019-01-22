package userValidator

import (
	"github.com/keidrun/boilerplate-gorilla-mux-for-rest-api-with-jwt/models"

	validator "gopkg.in/go-playground/validator.v9"
)

type UpdateUserRequest struct {
	Email    string            `validate:"omitempty,email"`
	Password string            `validate:"omitempty,min=6,max=30"`
	Name     string            `validate:"omitempty,min=3,max=20"`
	Age      models.NullInt64  `validate:"omitempty,gte=0"`
	Gender   models.NullString `validate:"omitempty,oneof=male female"`
}

func (v Validator) ValidateUpdateUserRequest(user models.User) []validator.FieldError {
	var target UpdateUserRequest
	target.Email = user.Email
	target.Password = user.Password
	target.Name = user.Name
	target.Age = user.Age
	target.Gender = user.Gender

	validate := NewValidator()
	err := validate.Struct(target)
	if err != nil {
		verrs := err.(validator.ValidationErrors)
		return verrs
	}
	return []validator.FieldError{}
}
