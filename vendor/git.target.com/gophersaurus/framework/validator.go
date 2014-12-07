package gophersauras

import (
	"errors"
	"reflect"
	"regexp"

	valLib "git.target.com/gophersaurus/gophersaurus/vendor/github.com/asaskevich/govalidator"
	valSys "git.target.com/gophersaurus/gophersaurus/vendor/gopkg.in/validator.v2"
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

// Pattern Validator gives the user the ability to create a validation function for a
// specific regular expression. This will allow developers to pair specific regular
// expressions to specific purposes.
func NewPatternValidator(pattern string) *patternValidator {
	return &patternValidator{regexp.MustCompile(pattern)}
}

type patternValidator struct {
	pattern *regexp.Regexp
}

func (p *patternValidator) Validate(v interface{}, param string) error {
	s := reflect.ValueOf(v)
	if s.Kind() != reflect.String {
		return errors.New("value not string")
	}
	if !p.pattern.MatchString(s.String()) {
		return errors.New("pattern not matched")
	}
	return nil
}
