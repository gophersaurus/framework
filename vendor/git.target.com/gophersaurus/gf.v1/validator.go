package gf

import (
	"errors"
	"net/http"
	"reflect"
	"regexp"

	valLib "git.target.com/gophersaurus/gophersaurus/vendor/github.com/asaskevich/govalidator"
	valSys "git.target.com/gophersaurus/gophersaurus/vendor/gopkg.in/validator.v2"
)

var Validator = valSys.NewValidator()

func Validate(v interface{}) error {
	return Validator.Validate(v)
}

func ShorthandValTag(name, tags string) {
	Validator.SetValidationFunc(name, func(v interface{}, param string) error {
		return Validator.Valid(v, tags)
	})
}

func ApplyLibValidator(name string, valFn valLib.Validator) {
	val := newLibValidator(name, valFn)
	ApplyErrorCode(val.Err, http.StatusBadRequest)
	Validator.SetValidationFunc(name, val.Validate)
}

func init() {
	Validator.SetTag("val")
	for name, fn := range valLib.TagMap {
		ApplyLibValidator(name, fn)
	}
	Validator.SetValidationFunc("email", email)
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

func (p *patternValidator) IsValid(v string) bool {
	return p.pattern.MatchString(v)
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
