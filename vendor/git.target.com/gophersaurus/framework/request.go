package gophersauras

import (
	"encoding/json"
	"errors"
	"net/http"

	"../../../github.com/gorilla/mux"
	"../../../gopkg.in/mgo.v2/bson"
)

type Query map[string][]string

type Request struct {
	Req       *http.Request
	Vars      map[string]string
	Query     Query
	SessionId bson.ObjectId
	Body      string
}

func NewRequest(req *http.Request) *Request {
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

	return &Request{
		Req:       req,
		Body:      string(jsonBody),
		SessionId: sessionId,

		// get the url path variables
		Vars: mux.Vars(req),

		// get url query parameters
		Query: (map[string][]string)(req.URL.Query()),
	}
}

func (r *Request) HasSession() bool {
	return r.SessionId != ""
}

func (r *Request) HasBody() bool {
	return len(r.Body) > 0
}

func (r *Request) ReadBody(obj interface{}) error {
	// Unmarshal JSON data into the cartwheelOffer object.
	err := json.Unmarshal([]byte(r.Body), obj)

	// Return an error if unable to unmarshal the cartwheelOffer ruest.
	if err != nil {
		return errors.New(InvalidJson)
	}
	return nil
}
