package models

import (
	gf "../../vendor/git.target.com/gophersaurus/framework"
	"../../vendor/gopkg.in/mgo.v2/bson"
)

type User struct {
	Id    bson.ObjectId `json:"_id" bson:"_id"`
	Email string        `json:"email" bson:"email"`
}

// NewUser creates an anonymous user.
func NewUser() *User {
	return &User{
		Id: bson.NewObjectId(),
	}
}

func (u *User) Find(key string, value interface{}) bool {
	err := gf.Mgo.C("users").Find(bson.M{key: value}).One(u)
	return err == nil // TODO: errors should really be checked for not found vs an actual error
}

func (u *User) Save() bool {
	_, err := gf.Mgo.C("users").UpsertId(u.Id, u)
	return err == nil // TODO: errors should really be checked for not found vs an actual error
}

func (u *User) Delete() bool {
	err := gf.Mgo.C("users").RemoveId(u.Id)
	return err == nil // TODO: errors should really be checked for not found vs an actual error
}
