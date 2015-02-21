package controllers

import (
	"errors"

	"git.target.com/gophersaurus/gf.v1"
	"git.target.com/gophersaurus/gophersaurus/app/models"
)

var Users = gf.NewResourceController("user_id",
	models.NewUser,
	models.FindAllUsers,
	func(req gf.Requester) (string, error) {
		sid := req.Header().Get("Session-Id")
		if len(sid) == 0 {
			return "", errors.New(gf.MissingSession)
		}
		session := models.NewSession()
		if err := session.FindByID(sid); err != nil {
			return "", err
		}
		return session.UserID.Hex(), nil
	})
