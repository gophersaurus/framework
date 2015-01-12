package models

import (
	"encoding/json"
	"errors"

	"git.target.com/gophersaurus/framework"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id        bson.ObjectId `json:"_id" bson:"_id"`
	Email     string        `json:"email,omitempty" bson:"email,omitempty" val:"email"`
	FirstName string        `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string        `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

// NewUser creates an anonymous user.
func NewUser() gf.Model {
	return &User{
		Id: bson.NewObjectId(),
	}
}

func (u *User) IdLabel() string {
	return "_id"
}

func (u *User) GetIdFrom(req gf.Request) (interface{}, error) {
	return gf.ObjectId(req)
}

func (u *User) SetId(id interface{}) error {
	objId, ok := id.(bson.ObjectId)
	if !ok {
		return errors.New(gf.InvalidId)
	}
	u.Id = objId
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

func (u *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(u)
}

func (u *User) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, u)
}
