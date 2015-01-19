package models

import (
	"errors"

	"git.target.com/gophersaurus/gf.v1"
	"gopkg.in/mgo.v2/bson"
)

type Address struct {
	user   *User
	Id     string `json:"id" bson:"id"`
	Street string `json:"street" bson:"street"`
	City   string `json:"city" bson:"city"`
	State  string `json:"state" bson:"state"`
	Zip    string `json:"zip" bson:"zip"`
}

func NewAddress() gf.Model {
	return &Address{Id: bson.NewObjectId().Hex()}
}

func (a *Address) SetId(id string) error {
	a.Id = id
	return nil
}

func (a *Address) BelongsTo(owner gf.Model) error {
	user, ok := owner.(*User)
	if !ok {
		return errors.New("invalid parent type")
	}
	a.user = user
	return nil
}

func (a *Address) FindById(id string) error {
	index, err := a.find(id)
	if err != nil {
		return err
	}
	if index < 0 {
		return errors.New("not found")
	}
	user := a.user
	*a = a.user.Addresses[index]
	a.user = user
	return nil
}

func (a *Address) Save() error {
	index, err := a.find(a.Id)
	if err != nil {
		return err
	}
	if index < 0 {
		a.user.Addresses = append(a.user.Addresses, *a)
	} else {
		a.user.Addresses[index] = *a
	}
	a.user.Save()
	return nil
}

func (a *Address) Delete() error {
	index, err := a.find(a.Id)
	if err != nil {
		return err
	}
	if index < 0 {
		return errors.New("not found")
	}
	addresses := []Address{}
	if index > 0 {
		addresses = append(addresses, a.user.Addresses[:index]...)
	}
	if index+1 < len(a.user.Addresses) {
		addresses = append(addresses, a.user.Addresses[index+1:]...)
	}
	a.user.Addresses = addresses
	a.user.Save()
	return nil
}

func (a *Address) Validate() error {
	return gf.Validate(a)
}

func (a *Address) find(id string) (int, error) {
	if a.user == nil {
		return 0, errors.New("parent not set")
	}
	for index, address := range a.user.Addresses {
		if address.Id == id {
			return index, nil
		}
	}
	return -1, nil
}
