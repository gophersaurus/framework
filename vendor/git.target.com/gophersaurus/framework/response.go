package gophersauras

import (
	"fmt"
	"net/http"
)

type Response struct {
	W      http.ResponseWriter
	status int
	body   []interface{}
}

func NewResponse(w http.ResponseWriter) *Response {
	return &Response{
		W:      w,
		status: http.StatusOK,
	}
}

func (r Response) RespondWithJSON(json map[string]string) {
	renderer.JSON(r.W, r.status, json)
}

func (r *Response) Respond() {
	if len(r.body) == 0 {
		// if no body present, respond with status code and headers only
		r.W.WriteHeader(r.status)
	} else {
		var body interface{}
		body = r.body
		if len(r.body) == 1 {
			// if only one element in body, send only that element not wrapped in an array
			body = r.body[0]
		}
		renderer.JSON(r.W, r.status, body)
	}
}

func (r *Response) RespondWithErr(err error) {

	// get the error message string
	message := err.Error()

	// build the message as json
	body := fmt.Sprintf("{\"error\": \"%v\"}", message)

	// look up the http code for the given error message, default to 500
	code, ok := errorMap[message]
	if !ok {
		if r.status != http.StatusOK {
			code = r.status
		} else {
			code = http.StatusInternalServerError
		}
	}

	// send error
	http.Error(r.W, body, code)
}

func (r *Response) HttpStatus(code int) {
	r.status = code
}

func (r *Response) Body(obj interface{}) {
	r.body = []interface{}{obj} // TODO: the '{}{obj}' seems very strange... are we doing what we think we are doing here?
}

func (r *Response) AppendBody(objects ...interface{}) {
	r.body = append(r.body, objects...)
}

func (r *Response) Header(key, value string) {
	r.W.Header().Set(key, value)
}

func (r *Response) FlushHeaders() { // TODO: we should revisit this function... is there another way to do this???
	headers := r.W.Header()

	// iterate through the headers and delete each
	for key, _ := range headers {
		delete(headers, key)
	}
}
