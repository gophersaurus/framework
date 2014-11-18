package app

import (
	gf "../vendor/git.target.com/gospot/framework"
	c "./controllers"
)

// register takes a Router and registers route paths to controller methods.
func register(r *gf.Router) {

	// Bind the HTTP GET path "/home" to the HelloWorldController Index() method.
	r.Get("/home", c.Home.Index)

	// Bind the HTTP POST path "/home" to the HelloWorldController Store() method.
	r.Post("/home", c.Home.Store)

	// Serve static files.
	r.Static("/", "public/")
}
