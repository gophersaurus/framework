package controllers

import "git.target.com/gophersaurus/gophersaurus/app/models"

var Addresses = Users.Extend("address_id", models.NewAddress, models.FindAllAddressesForUser)
