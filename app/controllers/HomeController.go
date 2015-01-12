package controllers

import (
	"strconv"

	"git.target.com/gophersaurus/framework"

	weather "git.target.com/gophersaurus/gophersaurus/app/services/openweathermap/data/v25"
)

// HomeController contains controller logic for home.
type HomeController struct{}

var Home = &HomeController{}

// Index handles a "/home" GET request for the HomeController.
func (h *HomeController) Index(resp gf.Response, req gf.Request) {

	w, err := weather.Find("Minneapolis", "us")

	if err != nil {

		// Add json body data.
		resp.RespondWithJSON(map[string]string{
			"hello Minneapolis": "Sorry, no weather report today. :( ",
			"error":             err.Error(),
		})

		return
	}

	// Add json body data.
	resp.RespondWithJSON(map[string]string{
		"Hello":       "Hi there, here is what's happening outside!",
		"location":    w.Name,
		"averageTemp": strconv.FormatFloat(w.Main.Temp, 'f', 1, 64),
		"highTemp":    strconv.FormatFloat(w.Main.TempMax, 'f', 1, 64),
		"lowTemp":     strconv.FormatFloat(w.Main.TempMin, 'f', 1, 64),
		"sunrise":     strconv.Itoa(w.Sys.Sunrise),
		"sunset":      strconv.Itoa(w.Sys.Sunset),
	})

}
