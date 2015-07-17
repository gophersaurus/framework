package controllers

import (
	"github.com/gophersaurus/gf.v1/http"

	weather "github.com/gophersaurus/framework/app/services/api.openweathermap.org/data/2.5"
)

// WeatherController contains controller logic for the weather endpoint.
type WeatherController struct{}

// Weather is a WeatherController.
var Weather = &WeatherController{}

// Show handles a "/weather/:city" GET request for a WeatherController.
func (wc *WeatherController) Show(resp http.Responder, req *http.Request) {

	// get the city as a parameter
	city := req.Param("city")

	// try some basic input checking
	if len(city) < 3 {
		resp.WriteErrs(req, http.InvalidInput, "not a valid city name")
		return
	}

	// use the weather service to get the weather
	w, err := weather.Find(city, "us")

	// check for errors
	if err != nil {

		// write a response
		resp.WriteErrs(req, "Sorry, no weather report today...", err.Error())
		return
	}

	// write the weather data
	resp.Write(req, w)
}
