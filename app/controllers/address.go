package controllers

import "git.target.com/gophersaurus/gophersaurus/app/models"

var Addresses = Users.Extend(models.NewAddress())
