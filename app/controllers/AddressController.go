package controllers

import (
	"git.target.com/gophersaurus/gf.v1"
	"git.target.com/gophersaurus/gophersaurus/app/models"
)

var Addresses = Users.Extend(gf.GetPathIDFunc("address_id"), models.NewAddress, models.FindAllAddressesForUser)
