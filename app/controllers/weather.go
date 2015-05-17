package controllers

import (
	"github.com/gophersaurus/gf.v1"

	weather "github.com/gophersaurus/framework/app/services/api.openweathermap.org/data/2.5"
)

// WeatherController contains controller logic for the weather endpoint.
type WeatherController struct{}

var Weather = &WeatherController{}

// Show handles a "/weather/:city" GET request for the WeatherController.
func (wc *WeatherController) Show(resp gf.Responder, req *gf.Request) {

	city := req.Param("city")

	if len(city) < 3 {
		resp.WriteErrs(gf.InvalidInput, "not a valid city name")
		return
	}

	w, err := weather.Find(city, "us")

	if err != nil {

		// Add json body data.
		resp.WriteJSON(map[string]string{
			"hello " + city: "Sorry, no weather report today. :( ",
			"error":         err.Error(),
		})

		return
	}

	// Add json body data.
	resp.Write(req, w)

}
