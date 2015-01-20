package middleware

import (
	"net/http"

	"git.target.com/gophersaurus/gf.v1"
	"git.target.com/gophersaurus/gophersaurus/app/models"
)

var SessionAdmin = NewSessionAdminMiddleware("Session-Id", "admin")

type SessionAdminMiddleware struct {
	SessionIDlabel string
	AdminRole      string
}

func NewSessionAdminMiddleware(sessionIDlabel, adminRole string) *SessionAdminMiddleware {
	return &SessionAdminMiddleware{sessionIDlabel, adminRole}
}

func (s *SessionAdminMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	errStr := w.Header().Get("Error")
	if len(errStr) > 0 {
		next(w, r)
		return
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
	next(w, r)
}
