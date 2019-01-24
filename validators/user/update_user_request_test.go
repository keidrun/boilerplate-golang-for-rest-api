package uservalidator

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/keidrun/boilerplate-gorilla-mux-for-rest-api-with-jwt/models"
)

type ValidateUpdateUserRequestTestCase struct {
	user     models.User
	expected []string
}

var validateUpdateUserRequestTestCases = []ValidateUpdateUserRequestTestCase{
	// OK
	{
		user:     models.User{Email: "testuser@example.com", Password: "testpassword", Name: "testuser", Age: models.NullInt64{NullInt64: sql.NullInt64{Int64: 20, Valid: true}}, Gender: models.NullString{NullString: sql.NullString{String: "male", Valid: true}}},
		expected: []string{},
	},
	{
		user:     models.User{},
		expected: []string{},
	},
	{
		user:     models.User{Email: "", Password: "", Name: "", Age: models.NullInt64{NullInt64: sql.NullInt64{Int64: 0, Valid: false}}, Gender: models.NullString{NullString: sql.NullString{String: "", Valid: false}}},
		expected: []string{},
	},
	// Email Error
	{
		user:     models.User{Email: "testuser#example.com", Password: "testpassword", Name: "testuser", Age: models.NullInt64{NullInt64: sql.NullInt64{Int64: 20, Valid: true}}, Gender: models.NullString{NullString: sql.NullString{String: "male", Valid: true}}},
		expected: []string{"Key: 'UpdateUserRequest.Email' Error:Field validation for 'Email' failed on the 'email' tag"},
	},
	// Pssword Error
	{
		user:     models.User{Email: "testuser@example.com", Password: "testu", Name: "testuser", Age: models.NullInt64{NullInt64: sql.NullInt64{Int64: 20, Valid: true}}, Gender: models.NullString{NullString: sql.NullString{String: "male", Valid: true}}},
		expected: []string{"Key: 'UpdateUserRequest.Password' Error:Field validation for 'Password' failed on the 'min' tag"},
	},
	{
		user:     models.User{Email: "testuser@example.com", Password: "1234567890123456789012345678901", Name: "testuser", Age: models.NullInt64{NullInt64: sql.NullInt64{Int64: 20, Valid: true}}, Gender: models.NullString{NullString: sql.NullString{String: "male", Valid: true}}},
		expected: []string{"Key: 'UpdateUserRequest.Password' Error:Field validation for 'Password' failed on the 'max' tag"},
	},
	// Name Error
	{
		user:     models.User{Email: "testuser@example.com", Password: "testpassword", Name: "te", Age: models.NullInt64{NullInt64: sql.NullInt64{Int64: 20, Valid: true}}, Gender: models.NullString{NullString: sql.NullString{String: "male", Valid: true}}},
		expected: []string{"Key: 'UpdateUserRequest.Name' Error:Field validation for 'Name' failed on the 'min' tag"},
	},
	{
		user:     models.User{Email: "testuser@example.com", Password: "testpassword", Name: "123456789012345678901", Age: models.NullInt64{NullInt64: sql.NullInt64{Int64: 20, Valid: true}}, Gender: models.NullString{NullString: sql.NullString{String: "male", Valid: true}}},
		expected: []string{"Key: 'UpdateUserRequest.Name' Error:Field validation for 'Name' failed on the 'max' tag"},
	},
	// Age Error
	{
		user:     models.User{Email: "testuser@example.com", Password: "testpassword", Name: "testuser", Age: models.NullInt64{NullInt64: sql.NullInt64{Int64: -1, Valid: true}}, Gender: models.NullString{NullString: sql.NullString{String: "male", Valid: true}}},
		expected: []string{"Key: 'UpdateUserRequest.Age' Error:Field validation for 'Age' failed on the 'gte' tag"},
	},
	// Gender Error
	{
		user:     models.User{Email: "testuser@example.com", Password: "testpassword", Name: "testuser", Age: models.NullInt64{NullInt64: sql.NullInt64{Int64: 20, Valid: true}}, Gender: models.NullString{NullString: sql.NullString{String: "superman", Valid: true}}},
		expected: []string{"Key: 'UpdateUserRequest.Gender' Error:Field validation for 'Gender' failed on the 'oneof' tag"},
	},
}

func TestValidateUpdateUserRequest(t *testing.T) {
	for index, test := range validateUpdateUserRequestTestCases {
		user := test.user
		validator := Validator{}
		actual := validator.ValidateUpdateUserRequest(user)

		if len(actual) != len(test.expected) {
			t.Errorf("[%d] Expected %d, got %d", index+1, len(test.expected), len(actual))
		}

		if reflect.DeepEqual(actual, test.expected) == false {
			t.Errorf("[%d] Expected \"%s\", got \"%s\"\n", index+1, test.expected, actual)
		}
	}
}
