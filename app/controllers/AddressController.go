package controllers

import (
	"git.target.com/gophersaurus/gophersaurus/app/models"
	"git.target.com/gophersaurus/gophersaurus/app/repos"
)

var Addresses = Users.Extend("address_id", models.NewAddress, repos.FindAllAddressesForUser)
