package framework

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"

	"../../../github.com/asaskevich/govalidator"
)

var Validate = newValidationMgr()

type Validator func(obj interface{}) error

type tag string

type validationMgr struct {
	library map[tag]Validator
}

func (m *validationMgr) RegisterTag(t string, v Validator) bool {
	_, exists := m.library[tag(t)]
	if exists {
		return false
	}
	m.library[tag(t)] = v
	return true
}

func (m *validationMgr) BuildFunc(obj interface{}, addl ...Validator) Validator {
	c := compositeValidator(addl)
	typeVal := newStructValidator(reflect.TypeOf(obj), m.library)
	if len(typeVal.fields) > 0 {
		c = append(c, typeVal.validate)
	}
	return (&c).validate
}

func newValidationMgr() *validationMgr {
	v := validationMgr{map[tag]Validator{}}
	for key, val := range govalidator.TagMap {
		err := errors.New("not " + key)
		ApplyErrorCode(err, http.StatusBadRequest)
		govalObj := &goval{err, val}
		v.RegisterTag(key, govalObj.val)
	}

	passwordPattern := "[a-zA-Z][a-zA-Z0-9]{7,19}$"
	v.RegisterTag("password", func(obj interface{}) error {
		if !govalidator.Matches(fmt.Sprintf("%v", obj), passwordPattern) {
			return errors.New(InvalidPassword)
		}
		return nil
	})
	return &v
}

type goval struct {
	err error
	fn  govalidator.Validator
}

func (g *goval) val(obj interface{}) error {
	if !g.fn(fmt.Sprintf("%v", obj)) {
		return g.err
	}
	return nil
}

type structValidator struct {
	typeOf reflect.Type
	fields []structFieldValidator
}

func newStructValidator(typeOf reflect.Type, library map[tag]Validator) *structValidator {
	if typeOf == nil || typeOf.Kind() != reflect.Struct {
		log.Fatal(errors.New("Cannot build validator for type other than struct"))
	}
	fields := []structFieldValidator{}
	size := typeOf.NumField()
	for i := 0; i < size; i++ {
		field := typeOf.Field(i)
		valList := []Validator{}
		tags := strings.Split(field.Tag.Get("validate"), ",")
		if len(tags) > 0 {
			for _, t := range tags {
				valFn, exists := library[tag(t)]
				if exists {
					valList = append(valList, valFn)
				}
			}
		}
		fieldType := field.Type
		if fieldType.Kind() == reflect.Struct {
			fieldTypeVal := newStructValidator(fieldType, library)
			if len(fieldTypeVal.fields) > 0 {
				valList = append(valList, fieldTypeVal.validate)
			}
		}
		if len(valList) > 0 {
			fields = append(fields, structFieldValidator{i, valList})
		}
	}
	return (&structValidator{typeOf, fields})
}

func (s *structValidator) validate(obj interface{}) error {
	value := reflect.ValueOf(obj)
	typeOf := value.Type()
	if typeOf != s.typeOf {
		log.Fatal(errors.New(fmt.Sprintf("Cannot validate instance of %v as type %v", typeOf.Name(), s.typeOf.Name())))
	}
	for _, field := range s.fields {
		fieldVal := value.Field(field.index)
		err := field.validators.validate(fieldVal.Interface())
		if err != nil {
			return err
		}
	}
	return nil
}

type structFieldValidator struct {
	index      int
	validators compositeValidator
}

type compositeValidator []Validator

func (c *compositeValidator) validate(obj interface{}) error {
	for _, val := range *c {
		err := val(obj)
		if err != nil {
			return err
		}
	}
	return nil
}
