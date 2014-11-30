package app

import (
	gf "../vendor/git.target.com/gophersaurus/framework"
	c "./controllers"
)

// register takes a Router and registers route paths to controller methods.
func register(r *gf.Router) {

	// Register the HTTP GET pattern "/home" to the HomeController Index() method.
	r.Get("/home", c.Home.Index)

	// Serve static files.
	r.Static("/", "public/")
}
