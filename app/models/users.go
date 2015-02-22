package models

import (
	"git.target.com/gophersaurus/gf.v1"
	"gopkg.in/mgo.v2/bson"
)

// FindAll gets all users from the database
func FindAllUsers() ([]gf.Model, error) {
	return fetchUsers(bson.M{})
}

func FindAllUsersByIds(ids ...bson.ObjectId) ([]gf.Model, error) {
	return fetchUsers(bson.M{
		"id": bson.M{
			"$in": ids,
		},
	})
}

func fetchUsers(query interface{}) ([]gf.Model, error) {
	var users []User
	err := dba.MGO("test").C("testUsers").Find(query).All(&users)
	if err != nil {
		return nil, err
	}
	out := []gf.Model{}
	for _, user := range users {
		temp := user // needed for pointer
		out = append(out, &temp)
	}
	return out, nil
}
