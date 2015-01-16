package models

import (
	"git.target.com/gophersaurus/gf.v1"
	"gopkg.in/mgo.v2/bson"
)

type User struct { // embedded object for marshalling and unmarshalling
	Id        bson.ObjectId `json:"_id" bson:"_id"`
	Email     string        `json:"email,omitempty" bson:"email,omitempty" val:"email"`
	FirstName string        `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string        `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Addresses []Address     `json:"addresses,omitempty" bson:"addresses,omitempty"`
}

// NewUser creates an anonymous user.
func NewUser() gf.Model {
	return &User{
		Id: bson.NewObjectId(),
	}
}

func (u *User) SetId(id string) error {
	bsonId, err := gf.StringToBsonId(id)
	u.Id = bsonId
	return err
}

func (u *User) FindById(id string) error {
	bsonId, err := gf.StringToBsonId(id)
	if err != nil {
		return err
	}
	return gf.Mgo.C("testUsers").FindId(bsonId).One(u)
}

func (u *User) Save() error {
	_, err := gf.Mgo.C("testUsers").UpsertId(u.Id, u)
	return err
}

func (u *User) Delete() error {
	return gf.Mgo.C("testUsers").RemoveId(u.Id)
}

func (u *User) Validate() error {
	return gf.Validate(u)
}
