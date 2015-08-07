package controllers

import (
	"github.com/gophersaurus/gf.v1/http"

	weather "github.com/gophersaurus/framework/app/services/api.openweathermap.org/data/2.5"
)

// Weather is a weather controller for cities.
var Weather = struct {
	Show func(resp http.Responder, req *http.Request)
}{
	Show: func(resp http.Responder, req *http.Request) {

		city := req.Param("city")

		// input checking
		if len(city) < 3 {
			resp.WriteErrs(req, http.InvalidInput, "not a valid city name")
			return
		}

		w, err := weather.Find(city, "us")
		if err != nil {
			resp.WriteErrs(req, "Sorry, no weather report today...", err.Error())
			return
		}

		// write response
		resp.Write(req, w)
	},
}
