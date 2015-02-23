package models

import (
	"errors"

	"git.target.com/gophersaurus/gf.v1"
	"gopkg.in/mgo.v2/bson"
)

type User struct { // embedded object for marshalling and unmarshalling
	Id        bson.ObjectId `json:"_id" bson:"_id"`
	Role      string        `json:"role,omitempty" bson:"role,omitempty"`
	Email     string        `json:"email,omitempty" bson:"email,omitempty" val:"email"`
	FirstName string        `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string        `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Addresses []Address     `json:"addresses,omitempty" bson:"addresses,omitempty"`
}

// NewUser creates an anonymous user.
func NewUser() *User {
	return &User{
		Id: bson.NewObjectId(),
	}
}

func (u *User) NewModel() gf.Model {
	return &User{
		Id: bson.NewObjectId(),
	}
}

func (u *User) PathID() string {
	return "user_id"
}

func (u *User) SetID(id string) error {
	bsonId, err := gf.BSONID(id)
	u.Id = bsonId
	return err
}

func (u *User) FindByID(id string) error {
	bsonId, err := gf.BSONID(id)
	if err != nil {
		return err
	}
	return dba.MGO("test").C("testUsers").FindId(bsonId).One(u)
}

func (u *User) FindAllByOwner(gf.Model) ([]gf.Model, error) {
	return nil, errors.New("user do not have an owner")
}

func (u *User) FindAll() ([]gf.Model, error) {

	// Get all the users.
	var users []User
	if err := dba.MGO("test").C("testUsers").Find(bson.M{}).All(&users); err != nil {
		return nil, err
	}

	// Unfortunately a []struct that implements []gf.Model is not compatible.
	// This is because of how interface memory is mapped in golang. It sucks.
	// More here: https://groups.google.com/forum/#!topic/golang-nuts/Il-tO1xtAyE
	models := make([]gf.Model, len(users))
	for i, v := range users {
		models[i] = gf.Model(&v)
	}

	return models, nil
}

func (u *User) Save() error {
	_, err := dba.MGO("test").C("testUsers").UpsertId(u.Id, u)
	return err
}

func (u *User) Delete() error {
	return dba.MGO("test").C("testUsers").RemoveId(u.Id)
}

func (u *User) Validate() error {
	return gf.Validate(u)
}

func (u *User) BelongsTo(owner gf.Model) error {
	return errors.New("user cannot belong to any other item")
}
