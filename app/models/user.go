package models

import (
	"git.target.com/gophersaurus/gf.v1"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id    bson.ObjectId `json:"_id" bson:"_id"`
	Email string        `json:"email" bson:"email" val:"email"`
}

// NewUser creates an anonymous user.
func NewUser() *User {
	return &User{
		Id: bson.NewObjectId(),
	}
}

func (u *User) Apply(patch gf.Patch) error {
	// TODO -- best way to apply subset of properties to object
	return nil
}

func (u *User) Find(key string, value interface{}) error {
	return gf.Mgo.C("testUsers").Find(bson.M{key: value}).One(u)
}

func (u *User) Save() error {
	_, err := gf.Mgo.C("testUsers").UpsertId(u.Id, u)
	return err
}

func (u *User) Delete() error {
	return gf.Mgo.C("testUsers").RemoveId(u.Id)
}
