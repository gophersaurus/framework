package controllers

import (
	"errors"

	"git.target.com/gophersaurus/gf.v1"
	"git.target.com/gophersaurus/gophersaurus/app/models"
)

// Users is a ResourceController.
var Users = gf.NewResourceController(models.NewUser(), func(req *gf.Request) (string, error) {

	// Get the session id from the header.
	if sid := req.Header().Get("Session-Id"); len(sid) > 0 {

		// Create a new session to work with.
		session := models.NewSession()

		// Find session by sid.
		if err := session.FindByID(sid); err != nil {
			return "", err
		}

		// Convert from user id from bson.ObjectId to string.
		return session.UserID.Hex(), nil
	}

	return "", errors.New(gf.MissingSession)
})
