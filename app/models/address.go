package models

import (
	"errors"

	"git.target.com/gophersaurus/gf.v1"
)

type Address struct {
	user   *User
	Id     string `json:"id" bson:"id"`
	Street string `json:"street" bson:"street"`
	City   string `json:"city" bson:"city"`
	State  string `json:"state" bson:"state"`
	Zip    string `json:"zip" bson:"zip"`
}

func NewAddress() gf.ChildModel {
	return &Address{}
}

func (a *Address) SetId(id interface{}) error {
	objId, ok := id.(string)
	if !ok {
		return errors.New(gf.InvalidId)
	}
	a.Id = objId
	return nil
}

func (a *Address) SetParent(parent gf.Model) error {
	user, ok := parent.(*User)
	if !ok {
		return errors.New("invalid parent type")
	}
	a.user = user
	return nil
}

func (a *Address) FindById(id interface{}) error {
	// TODO
	return nil
}

func (a *Address) Save() error {
	// TODO
	return nil
}

func (a *Address) Delete() error {
	// TODO
	return nil
}

func (a *Address) find(id interface{}) (int, error) {
	if a.user == nil {
		return 0, errors.New("parent not set")
	}
	/*
		for index, address := range a.user.Addresses {

		}
	*/
	return -1, nil
}
