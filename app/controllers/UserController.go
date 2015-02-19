package controllers

import (
	"git.target.com/gophersaurus/gf.v1"
	"git.target.com/gophersaurus/gophersaurus/app/models"
)

var Users = gf.NewResourceController(gf.GetPathIDFunc("user_id"), models.NewUser, models.FindAllUsers)
