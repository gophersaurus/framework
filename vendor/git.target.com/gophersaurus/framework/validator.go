package gophersauras

import (
	"errors"
	"reflect"

	valLib "../../../github.com/asaskevich/govalidator"
	valSys "../../../gopkg.in/validator.v2"
)

func init() {
	Validator := valSys.NewValidator()
	Validator.SetTag("val")
	Validator.SetValidationFunc("email", email)
}

var Validator *valSys.Validator

func Validate(v interface{}) error {
	return Validator.Validate(v)
}

func email(v interface{}, param string) error {

	s := reflect.ValueOf(v)
	if s.Kind() != reflect.String {
		return errors.New("the email validator only validates strings")
	}

	if !valLib.IsEmail(s.String()) {
		return errors.New("not a valid email")
	}

	return nil
}

// passwordPattern := "[a-zA-Z][a-zA-Z0-9]{7,19}$"
