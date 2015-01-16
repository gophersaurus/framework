package controllers

import (
	"git.target.com/gophersaurus/gophersaurus/app/models"
	"git.target.com/gophersaurus/gophersaurus/app/repos"

	"git.target.com/gophersaurus/gf.v1"
)

var User = gf.NewResource("user_id", models.NewUser, repos.FindAllUsers)
