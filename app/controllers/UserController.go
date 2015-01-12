package controllers

import (
	"git.target.com/gophersaurus/gophersaurus/app/models"
	"git.target.com/gophersaurus/gophersaurus/app/repos"

	"git.target.com/gophersaurus/framework"
)

var User = gf.NewResource(repos.Users, models.NewUser)
