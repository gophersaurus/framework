package repos

import (
	"git.target.com/gophersaurus/gophersaurus/app/models"
	gf "git.target.com/gophersaurus/gophersaurus/vendor/git.target.com/gophersaurus/framework"
	"git.target.com/gophersaurus/gophersaurus/vendor/gopkg.in/mgo.v2/bson"
)

// FindAllUsers gets all users from the database
func FindAllUsers() ([]models.User, error) {
	var users []models.User
	err := gf.Mgo.C("users").Find(bson.M{}).All(&users)
	return users, err
}

// FindUsersById gets all users whose id is in the given list
func FindUsersById(ids ...string) ([]models.User, error) {
	query := bson.M{
		"id": bson.M{
			"$in": ids,
		},
	}
	var users []models.Users
	err := gf.Mgo.C("users").Find(query).All(&users)
	return users, err
}
