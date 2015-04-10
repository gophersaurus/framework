package middleware

import (
	"net/http"

	"github.com/gophersaurus/gf.v1"
	"github.com/gophersaurus/framework/app/models"
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
	sessionID, err := gf.BSONID(r.Header.Get(s.SessionIDlabel))
	if err != nil {
		w.Header().Set("Error", gf.MissingSession)
		next(w, r)
		return
	}

	session := models.NewSession()

	// Search for session by sessionId.
	if session_db_err := session.FindByID(sessionID.Hex()); session_db_err != nil {
		w.Header().Set("Error", gf.MissingSession)
		next(w, r)
		return
	}

	user := &models.User{}
	if user_db_err := user.FindByID(session.UserID.Hex()); user_db_err != nil {
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
