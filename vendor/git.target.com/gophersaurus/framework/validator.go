package gophersauras

import (
	"errors"
	"net/http"
	"reflect"
	"regexp"

	valLib "git.target.com/gophersaurus/gophersaurus/vendor/github.com/asaskevich/govalidator"
	valSys "git.target.com/gophersaurus/gophersaurus/vendor/gopkg.in/validator.v2"
)

func init() {
	Validator := valSys.NewValidator()
	Validator.SetTag("val")
	for name, fn := range valLib.TagMap {
		val := newLibValidator(name, fn)
		ApplyErrorCode(val.Err, http.StatusBadRequest)
		Validator.SetValidationFunc(name, val.Validate)
	}
	Validator.SetValidationFunc("email", email)
}

var Validator = valSys.NewValidator()

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

type libValidator struct {
	name  string
	Err   error
	valFn valLib.Validator
}

func newLibValidator(name string, valFn valLib.Validator) *libValidator {
	return &libValidator{name, errors.New("not " + name), valFn}
}

func (l *libValidator) Validate(v interface{}, param string) error {
	s := reflect.ValueOf(v)
	if s.Kind() != reflect.String {
		return errors.New(l.name + " validator only validates string values")
	}
	if !l.valFn(s.String()) {
		return l.Err
	}
	return nil
}
