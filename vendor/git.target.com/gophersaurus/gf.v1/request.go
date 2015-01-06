package gf

import (
	"encoding/json"
	"errors"
	"net/http"

	"git.target.com/gophersaurus/gophersaurus/vendor/github.com/gorilla/mux"
	"git.target.com/gophersaurus/gophersaurus/vendor/gopkg.in/mgo.v2/bson"
)

type Request interface {
	Request() *http.Request
	Var(name string) string
	Query(name string) []string
	SessionId() bson.ObjectId
	Body() string
	HasSession() bool
	HasBody() bool
	ReadBody(obj interface{}) error
}

type request struct {
	req       *http.Request
	vars      map[string]string
	query     map[string][]string
	sessionId bson.ObjectId
	body      string
}

func buildRequest(req *http.Request) Request {
	// Make a slice of bytes large enough for the JSON ruest body.
	jsonBody := make([]byte, req.ContentLength)

	// Read the json body bytes.
	req.Body.Read(jsonBody)

	// get the sessionId from the header and convert it to a
	sessionIdStr := req.Header.Get("Session-Id")
	var sessionId bson.ObjectId = ""
	if bson.IsObjectIdHex(sessionIdStr) {
		sessionId = bson.ObjectIdHex(sessionIdStr)
	} else {
		sessionId = ""
	}

	return &request{
		req:       req,
		body:      string(jsonBody),
		sessionId: sessionId,

		// get the url path variables
		vars: mux.Vars(req),

		// get url query parameters
		query: (map[string][]string)(req.URL.Query()),
	}
}

func (r *request) Request() *http.Request {
	return r.req
}

func (r *request) Var(name string) string {
	return r.vars[name]
}

func (r *request) Query(name string) []string {
	return r.query[name]
}

func (r *request) SessionId() bson.ObjectId {
	return r.sessionId
}

func (r *request) Body() string {
	return r.body
}

func (r *request) HasSession() bool {
	return r.sessionId != ""
}

func (r *request) HasBody() bool {
	return len(r.body) > 0
}

func (r *request) ReadBody(obj interface{}) error {
	// Unmarshal JSON data into the cartwheelOffer object.
	err := json.Unmarshal([]byte(r.body), obj)

	// Return an error if unable to unmarshal the cartwheelOffer ruest.
	if err != nil {
		return errors.New(InvalidJson)
	}
	return nil
}
