package middleware

import (
	"git.target.com/gophersaurus/gf.v1"
	"git.target.com/gophersaurus/gophersaurus/app/models"
)

var SessionUser = NewSessionUserMiddleware("Session-Id", "user_id")

type SessionUserMiddleware struct {
	SessionIDlabel string
	UserIDlabel    string
}

func NewSessionUserMiddleware(sessionIDlabel, userIDlabel string) *SessionUserMiddleware {
	return &SessionUserMiddleware{sessionIDlabel, userIDlabel}
}

func (s *SessionUserMiddleware) ServeHTTP(resp gf.Responder, req *gf.Request, next gf.HandlerFunc) {

	// check to ensure the presence of a session
	sessionID, err := gf.BSONID(req.Header().Get(s.SessionIDlabel))
	if err != nil {
		resp.WriteErrs(gf.MissingSession)
		return
	}

	session := models.NewSession()

	// Search for session by sessionId.
	if session_db_err := session.FindByID(sessionID.Hex()); session_db_err != nil {
		resp.WriteErrs(gf.MissingSession)
		return
	}

	/*
		if the user of the session is not the user in the path, pass error
	*/
	userVar := req.Param(s.UserIDlabel)
	if session.UserID.Hex() != userVar {
		resp.WriteErrs(gf.InvalidPermission)
		return
	}
	next(resp, req)
}
