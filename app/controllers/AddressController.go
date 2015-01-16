package controllers

import (
	"git.target.com/gophersaurus/gf.v1"
	"git.target.com/gophersaurus/gophersaurus/app/models"
	"git.target.com/gophersaurus/gophersaurus/app/repos"
)

var Addresses = gf.NewExtendedResource(Users, "address_id", models.NewAddress, repos.FindAllAddressesForUser)
