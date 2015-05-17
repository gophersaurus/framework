package controllers

import "github.com/gophersaurus/gf.v1"

// HomeController contains controller logic for home.
type HomeController struct{}

// Home is a HomeController.
var Home = &HomeController{}

// Index handles a "/" GET request for a HomeController.
func (h *HomeController) Index(resp gf.Responder, req *gf.Request) {

	// set the default HTTP scheme without SSL/TLS
	scheme := "http://"

	// check if we are serving SSL/TLS HTTP traffic
	if req.TLS != nil {
		scheme = "https://"
	}

	// define an anonymous result struct
	result := struct {
		Status     int    `json:"status" xml:"status"`
		Message    string `json:"message" xml:"message"`
		PublicPage string `json:"public_page" xml:"public_page"`
	}{
		200,
		"Welcome fellow gopher.",
		scheme + req.Host + "/public/index.html",
	}

	// write the result
	resp.Write(req, result)
}
