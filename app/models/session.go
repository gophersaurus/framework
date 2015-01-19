package models

import (
	"errors"
	"time"

	"git.target.com/gophersaurus/gf.v1"

	"gopkg.in/mgo.v2/bson"
)

type Session struct {
	ID      bson.ObjectId `json:"_id" bson:"_id"`
	UserID  bson.ObjectId `json:"userId" bson:"userId"`
	Expires int64         `json:"expires" bson:"expires"`
	User    *User         `json:"user,omitempty" bson:"omitempty"`
}

func NewSession() *Session {
	return &Session{
		ID:      bson.NewObjectId(),
		Expires: time.Now().Add(30 * 24 * time.Hour).Unix(),
	}
}

func (s *Session) IsExpired() bool {
	return time.Now().Unix() > s.Expires
}

func (s *Session) SetId(id string) error {
	bsonId, err := gf.StringToBsonID(id)
	s.ID = bsonId
	return err
}

func (s *Session) FindById(id string) error {
	bsonId, err := gf.StringToBsonID(id)
	if err != nil {
		return err
	}
	return gf.Mgo.C("testSessions").FindId(bsonId).One(s)
}

func (s *Session) Save() error {
	s.User = nil
	_, err := gf.Mgo.C("testSessions").UpsertId(s.ID, s)
	return err
}

func (s *Session) Delete() error {
	return gf.Mgo.C("testSessions").RemoveId(s.ID)
}

func (s *Session) Validate() error {
	return gf.Validate(s)
}

func (s *Session) BelongsTo(owner gf.Model) error {
	return errors.New("session cannot have owner")
}
