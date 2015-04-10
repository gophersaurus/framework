package controllers

import (
	"git.target.com/gophersaurus/gf.v1"

	weather "github.com/gophersaurus/framework/app/services/api.openweathermap.org/data/2.5"
)

// HomeController contains controller logic for home.
type HomeController struct{}

var Home = &HomeController{}

// Index handles a "/home" GET request for the HomeController.
func (h *HomeController) Index(resp gf.Responder, req *gf.Request) {

	w, err := weather.Find("minneapolis", "us")

	if err != nil {

		// Add json body data.
		resp.WriteJSON(map[string]string{
			"hello Minneapolis": "Sorry, no weather report today. :( ",
			"error":             err.Error(),
		})

		return
	}

	// Add json body data.
	resp.Write(req, w)

}
