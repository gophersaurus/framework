package middleware

import (
	"log"
	"net/http"

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

func (s *SessionUserMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
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
		w.Header().Set("Error", gf.InvalidPermission)
		next(w, r)
		return
	}
	next(w, r)
}
