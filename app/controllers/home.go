package controllers

import "github.com/gophersaurus/gf.v1/http"

// Home is a home controller for new users.
var Home = struct {
	Index func(resp http.Responder, req *http.Request)
}{
	Index: func(resp http.Responder, req *http.Request) {

		// define a result
		result := struct {
			Status        int    `json:"status" xml:"status"`
			Message       string `json:"message" xml:"message"`
			PublicPage    string `json:"public_page" xml:"public_page"`
			PublicAPIDocs string `json:"public_api_docs" xml:"public_api_docs"`
		}{
			200,
			"Welcome fellow gopher.",
			"http://" + req.Host + "/public",
			"http://" + req.Host + "/public/docs/api",
		}

		// write the result
		resp.Write(req, result)
	},
}
