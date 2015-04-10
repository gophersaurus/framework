package middleware

import (
	"github.com/gophersaurus/gf.v1"
	"github.com/gophersaurus/framework/app/models"
)

var SessionUserAdmin = NewSessionUserAdminMiddleware("Session-Id", "user_id", "admin")

type SessionUserAdminMiddleware struct {
	SessionIDlabel string
	UserIDlabel    string
	AdminRole      string
}

func NewSessionUserAdminMiddleware(sessionIDlabel, userIDlabel, adminRole string) *SessionUserAdminMiddleware {
	return &SessionUserAdminMiddleware{sessionIDlabel, userIDlabel, adminRole}
}

func (s *SessionUserAdminMiddleware) ServeHTTP(resp gf.Responder, req *gf.Request, next gf.HandlerFunc) {

	// check to ensure the presence of a session
	sessionID, err := gf.BSONID(req.Header.Get(s.SessionIDlabel))
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
		user := &models.User{}
		if user_db_err := user.FindByID(session.UserID.Hex()); user_db_err != nil {
			resp.WriteErrs(gf.MissingSession)
			return
		}
		if user.Role != s.AdminRole {
			resp.WriteErrs(gf.MissingSession)
			return
		}
	}

	next(resp, req)

}
