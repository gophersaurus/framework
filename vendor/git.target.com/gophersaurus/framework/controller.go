package gophersauras

import "net/http"

type Controller interface {
	Index(resp *Response, req *Request)
	Store(resp *Response, req *Request)
	Show(resp *Response, req *Request)
	Destroy(resp *Response, req *Request)
}

type Updateable interface {
	Controller
	Update(resp *Response, req *Request)
}

type Patchable interface {
	Controller
	Apply(resp *Response, req *Request)
}

type Patch map[string]interface{}

type Action func(resp *Response, req *Request)

func Handle(action Action) http.HandlerFunc {
	return action.ServeHTTP
}

func (a Action) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp := buildResponse(w)
	req := buildRequest(r)
	a(resp, req)
}
