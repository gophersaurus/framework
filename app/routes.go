package server

import (
	c "github.com/gophersaurus/framework/app/controllers"
	"github.com/gophersaurus/gf.v1"
)

// register takes a Router and registers route paths to controller methods.
func register(r *gf.Router) {

	// Register the HTTP GET pattern "/home" to the HomeController Index() method.
	r.GET("/", c.Home.Index)

	// Register the HTTP GET pattern "/work" to the WorkController Index() method.
	r.GET("/weather/:city/", c.Weather.Show)

}
