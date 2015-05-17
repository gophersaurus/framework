package controllers

import "github.com/gophersaurus/gf.v1"

// HomeController contains controller logic for home.
type HomeController struct{}

var Home = &HomeController{}

// Index handles a "/home" GET request for the HomeController.
func (h *HomeController) Index(resp gf.Responder, req *gf.Request) {

	// set the default HTTP scheme without SSL/TLS
	scheme := "http://"

	// check if we are serving SSL/TLS HTTP traffic
	if req.TLS != nil {
		scheme = "https://"
	}

	// define an anonymous result struct
	result := struct {
		Status     int
		Message    string
		StaticPage string
	}{
		200,
		"You have arrived.",
		scheme + req.Host + "/public/index.html",
	}

	// write the result
	resp.Write(req, result)
}
