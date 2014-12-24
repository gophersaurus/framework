package gf

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"git.target.com/gophersaurus/gophersaurus/vendor/github.com/gorilla/mux"
	"git.target.com/gophersaurus/gophersaurus/vendor/gopkg.in/mgo.v2/bson"
)

type Query map[string][]string

type Request struct {
	Req       *http.Request
	Vars      map[string]string
	Query     Query
	SessionId bson.ObjectId
	Body      string
}

func NewRequest(method, url string, body []byte) (*Request, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	return &Request{Req: req, Body: string(body)}, nil
}

func (r *Request) AddHeader(name, value string) {
	r.Req.Header.Add(name, value)
}

func (r *Request) SetHeader(name, value string) {
	r.Req.Header.Set(name, value)
}

func (r *Request) Send() ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(r.Req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	return bytes, err

}

func buildRequest(req *http.Request) *Request {
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
