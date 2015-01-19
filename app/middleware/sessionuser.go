package middleware

import (
	"log"
	"net/http"

	"git.target.com/gophersaurus/gf.v1"
	"git.target.com/gophersaurus/gophersaurus/app/models"
)

type SessionUserMiddleware struct {
	SessionIDlabel string
	UserIDlabel    string
	AdminRole      string
}

func NewSessionUserMiddleware(sessionIDlabel, userIDlabel, adminRole string) *SessionUserMiddleware {
	return &SessionUserMiddleware{sessionIDlabel, userIDlabel, adminRole}
}

func (s *SessionUserMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// Create a new Response.
	resp := gf.NewResponse(w)

	// Create a new Request.
	req, err := gf.NewRequest(r)
	if err != nil {
		log.Fatal(err)
	}

	// check to ensure the presence of a session
	sessionID, err := gf.StringToBsonID(r.Header.Get(s.SessionIDlabel))
	if err != nil {
		resp.RespondWithErr(gf.MissingSession)
		return
	}

	session := models.NewSession()

	// Search for session by sessionId.
	if session_db_err := session.FindById(sessionID.Hex()); session_db_err != nil {
		resp.RespondWithErr(gf.MissingSession)
		return
	}

	/*
		if the user of the session is not the user in the path, then
		the user of the session must have admin permission to edit
	*/
	userVar, _ := req.Var(s.UserIDlabel)
	if session.UserID.Hex() != userVar {
		// Search for user by session.
		user := &models.User{}
		if user_db_err := user.FindById(session.UserID.Hex()); user_db_err != nil {
			resp.RespondWithErr(gf.MissingUser)
			return
		}
		if user.Role != s.AdminRole {
			resp.RespondWithErr(gf.InvalidPermission)
			return
		}
	}
	next(w, r)
}
