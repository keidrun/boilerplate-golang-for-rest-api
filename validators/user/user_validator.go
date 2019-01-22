package userValidator

import (
	"database/sql/driver"
	"log"
	"reflect"

	"github.com/keidrun/boilerplate-gorilla-mux-for-rest-api-with-jwt/models"
	validator "gopkg.in/go-playground/validator.v9"
)

type Validator struct{}

func NewValidator() *validator.Validate {
	validate := validator.New()
	validate.RegisterCustomTypeFunc(validateValuer,
		models.NullBool{}, models.NullFloat64{}, models.NullInt64{}, models.NullString{})
	return validate
}

func validateValuer(field reflect.Value) interface{} {
	if valuer, ok := field.Interface().(driver.Valuer); ok {
		value, err := valuer.Value()
		if err != nil {
			log.Fatal(err)
		}
		return value
	}
	return nil
}
