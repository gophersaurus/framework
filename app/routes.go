package server

import (
	"git.target.com/gophersaurus/gf.v1"
	c "github.com/gophersaurus/framework/app/controllers"
)

// register takes a Router and registers route paths to controller methods.
func register(r *gf.Router) {

	// Register the HTTP GET pattern "/home" to the HomeController Index() method.
	r.GET("/home", c.Home.Index)

	// Register the HTTP GET pattern "/work" to the WorkController Index() method.
	r.GET("/work", c.Work.Index)

}
