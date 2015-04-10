package controllers

import "github.com/gophersaurus/framework/app/models"

var Addresses = Users.Extend(models.NewAddress())
