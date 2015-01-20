package middleware

import (
	"log"
	"net/http"

	"git.target.com/gophersaurus/gf.v1"
	"git.target.com/gophersaurus/gophersaurus/app/models"
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

func (s *SessionUserAdminMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	errStr := w.Header().Get("Error")
	if len(errStr) > 0 {
		next(w, r)
		return
	}

	// Create a new Request.
	req, err := gf.NewRequest(r)
	if err != nil {
		log.Fatal(err)
	}

	// check to ensure the presence of a session
	sessionID, err := gf.StringToBsonID(r.Header.Get(s.SessionIDlabel))
	if err != nil {
		w.Header().Set("Error", gf.MissingSession)
		next(w, r)
		return
	}

	session := models.NewSession()

	// Search for session by sessionId.
	if session_db_err := session.FindById(sessionID.Hex()); session_db_err != nil {
		w.Header().Set("Error", gf.MissingSession)
		next(w, r)
		return
	}

	/*
		if the user of the session is not the user in the path, pass error
	*/
	userVar, _ := req.Var(s.UserIDlabel)
	if session.UserID.Hex() != userVar {
		user := &models.User{}
		if user_db_err := user.FindById(session.UserID.Hex()); user_db_err != nil {
			w.Header().Set("Error", gf.MissingUser)
			next(w, r)
			return
		}
		if user.Role != s.AdminRole {
			w.Header().Set("Error", gf.InvalidPermission)
			next(w, r)
			return
		}
	}
	next(w, r)

}
