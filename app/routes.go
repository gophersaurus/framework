package app

import (
	"git.target.com/gophersaurus/gf.v1"
	c "git.target.com/gophersaurus/gophersaurus/app/controllers"
)

// register takes a Router and registers route paths to controller methods.
func register(r *gf.Router) {

	// Register the HTTP GET pattern "/home" to the HomeController Index() method.
	r.Get("/home", c.Home.Index)

	// Register the HTTP GET pattern "/work" to the WorkController Index() method.
	r.Get("/work", c.Work.Index)

	// Register the pattern "/user" with all methods in the User Controller
	r.Resource("/user", "user_id", c.Users)

	// Register the pattern "/user/{user_id}/address" with all methods in the Address Controller
	r.Resource("/user/{user_id}/address", "address_id", c.Addresses)

	// Serve static files.
	r.Static("/", "public/")
}
