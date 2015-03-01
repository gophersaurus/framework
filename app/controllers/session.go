package controllers

import (
	"fmt"

	"git.target.com/gophersaurus/gf.v1"
	"git.target.com/gophersaurus/gophersaurus/app/models"
)

var Sessions = &SessionController{}

type SessionController struct {
}

func (s *SessionController) Store(resp gf.Responder, req gf.Requester) {
	body := map[string]string{}
	err := req.Read(&body)
	if err != nil {
		resp.WriteErrs(err.Error())
		return
	}

	userIDstr, ok := body["user"]
	if !ok {
		resp.WriteErrs("missing user id")
		return
	}

	userID, err := gf.BSONID(userIDstr)
	if err != nil {
		resp.WriteErrs(err.Error())
		return
	}

	user := models.NewUser()
	err = user.FindByID(userID.Hex())
	if err != nil {
		resp.WriteErrs("Invalid User")
		return
	}

	session := models.NewSession()
	session.UserID = userID
	err = session.Save()
	if err != nil {
		resp.WriteErrs(err.Error())
		return
	}
	resp.Header("Session-Id", session.ID.Hex())
	resp.Header("Session-Expires", fmt.Sprintf("%v", session.Expires))
	resp.Do(req)
}

func (s *SessionController) Show(resp gf.Responder, req gf.Requester) {
	sessionID, err := gf.BSONID(req.Header().Get("Session-Id"))
	if err != nil {
		resp.WriteErrs(gf.MissingSession)
		return
	}

	session := models.NewSession()
	err = session.FindByID(sessionID.Hex())
	if err != nil {
		resp.WriteErrs(gf.MissingSession)
		return
	}

	user := &models.User{}
	err = user.FindByID(session.UserID.Hex())
	if err != nil {
		resp.WriteErrs("Invalid User")
		return
	}

	session.User = user

	resp.Read(session)
	resp.JSON()
}
