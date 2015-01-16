package controllers

import (
	"git.target.com/gophersaurus/gophersaurus/app/models"
	"git.target.com/gophersaurus/gophersaurus/app/repos"

	"git.target.com/gophersaurus/gf.v1"
)

var Users = gf.NewResource("user_id", models.NewUser, repos.FindAllUsers)
