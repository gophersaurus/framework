package gf

import "net/http"

type Response interface {
	RespondWithJSON(json map[string]string)
	Respond()
	RespondWithErr(err error)
	HttpStatus(code int)
	Body(obj interface{})
	AppendBody(objects ...interface{})
	Header(key, value string)
	FlushHeaders()
}

type response struct {
	w      http.ResponseWriter
	status int
	body   []interface{}
}

func buildResponse(w http.ResponseWriter) Response {
	return &response{
		w:      w,
		status: http.StatusOK,
	}
}

func (r *response) RespondWithJSON(json map[string]string) {
	renderer.JSON(r.w, r.status, json)
}

func (r *response) Respond() {
	if len(r.body) == 0 {
		// if no body present, respond with status code and headers only
		r.w.WriteHeader(r.status)
	} else {
		var body interface{}
		body = r.body
		if len(r.body) == 1 {
			// if only one element in body, send only that element not wrapped in an array
			body = r.body[0]
		}
		renderer.JSON(r.w, r.status, body)
	}
}

func (r *response) RespondWithErr(err error) {

	// get the error message string
	message := err.Error()

	// build the message as json
	body := map[string]string{"error": message}

	// look up the http code for the given error message, default to 500
	code, ok := errorMap[message]
	if !ok {
		if r.status != http.StatusOK {
			code = r.status
		} else {
			code = http.StatusInternalServerError
		}
	}
	r.status = code

	r.RespondWithJSON(body)
}

func (r *response) HttpStatus(code int) {
	r.status = code
}

func (r *response) Body(obj interface{}) {
	r.body = []interface{}{obj} // TODO: the '{}{obj}' seems very strange... are we doing what we think we are doing here?
}

func (r *response) AppendBody(objects ...interface{}) {
	r.body = append(r.body, objects...)
}

func (r *response) Header(key, value string) {
	r.w.Header().Set(key, value)
}

func (r *response) FlushHeaders() { // TODO: we should revisit this function... is there another way to do this???
	headers := r.w.Header()

	// iterate through the headers and delete each
	for key, _ := range headers {
		delete(headers, key)
	}
}
