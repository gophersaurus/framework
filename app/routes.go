package app

import (
	c "github.com/gophersaurus/framework/app/controllers"
	"github.com/gophersaurus/gf.v1/router"
)

// register takes a Router and registers route paths to controller methods.
func register(r router.Router) {

	// Register the HTTP GET URI "/home" to the Home controller Index method.
	r.GET("/", c.Home.Index)

	// Register the HTTP GET URI "/weather" to the Weather controller Show method.
	r.GET("/weather/:city", c.Weather.Show)

}
