package repos

import (
	"git.target.com/gophersaurus/gophersaurus/app/models"
	"git.target.com/gophersaurus/gf.v1"
	"gopkg.in/mgo.v2/bson"
)

// FindAllUsers gets all users from the database
func FindAllUsers() ([]models.User, error) {
	var users []models.User
	err := gf.Mgo.C("testUsers").Find(bson.M{}).All(&users)
	return users, err
}

// FindUsersById gets all users whose id is in the given list
func FindUsersById(ids ...string) ([]models.User, error) {
	query := bson.M{
		"id": bson.M{
			"$in": ids,
		},
	}
	var users []models.User
	err := gf.Mgo.C("testUsers").Find(query).All(&users)
	return users, err
}
