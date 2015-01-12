package repos

import (
	"strings"

	"git.target.com/gophersaurus/framework"
	"git.target.com/gophersaurus/gophersaurus/app/models"
	"gopkg.in/mgo.v2/bson"
)

var Users = &users{}

type users struct {
}

// FindAllUsers gets all users from the database
func (u *users) FindAll() ([]gf.Model, error) {
	return u.fetch(bson.M{})
}

func (u *users) GetQuery(req gf.Request) interface{} {
	return bson.M{
		"id": bson.M{
			"$in": strings.Split(req.Request().URL.Query().Get("ids"), ","),
		},
	}
}

// FindUsersById gets all users whose id is in the given list
func (u *users) Query(query interface{}) ([]gf.Model, error) {
	return u.fetch(query)
}

func (u *users) fetch(query interface{}) ([]gf.Model, error) {
	var users []models.User
	err := gf.Mgo.C("testUsers").Find(query).All(&users)
	if err != nil {
		return nil, err
	}
	out := make([]gf.Model, len(users))
	for index, user := range users {
		out[index] = &user
	}
	return out, nil
}
