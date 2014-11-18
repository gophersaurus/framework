package controllers

import (
	"strconv"

	gf "../../vendor/git.target.com/gophersaurus/framework"
	weather "../services/openweathermap/data/v25"

	"../validators/forms"
)

// HomeController contains controller logic for home.
type HomeController struct{}

var Home *HomeController

func init() {
	Home = &HomeController{}
}

// Index handles a "/home" GET request for the HomeController.
func (h *HomeController) Index(resp *gf.Response, req *gf.Request) {

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
		"hello Minneapolis":             "Here is what's happening outside!",
		"location":                      w.Name,
		"average temp (°C)":             strconv.FormatFloat(w.Main.Temp, 'f', 1, 64),
		"high temp (°C)":                strconv.FormatFloat(w.Main.TempMax, 'f', 1, 64),
		"low temp (°C)":                 strconv.FormatFloat(w.Main.TempMin, 'f', 1, 64),
		"next sunrise (unix timestamp)": strconv.Itoa(w.Sys.Sunrise),
		"next sunset (unix timestamp)":  strconv.Itoa(w.Sys.Sunset),
	})

}

func (s *HomeController) Store(resp *gf.Response, req *gf.Request) {
	err := forms.Login(req)
	if err != nil {
		resp.RespondWithErr(err)
		return
	}
	resp.Respond()
}
