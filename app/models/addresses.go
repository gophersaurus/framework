package models

import (
	"errors"
	"reflect"

	"git.target.com/gophersaurus/gf.v1"
)

func FindAllAddressesForUser(owner gf.Model) ([]gf.Model, error) {
	value := reflect.ValueOf(owner).Elem().Interface()
	user, ok := value.(User)
	if !ok {
		return nil, errors.New("invalid parent type")
	}
	out := []gf.Model{}
	for _, address := range user.Addresses {
		temp := address
		out = append(out, &temp)
	}
	return out, nil
}
