package app

import (
	"git.target.com/gophersaurus/gf.v1"
	c "git.target.com/gophersaurus/gophersaurus/app/controllers"
	"git.target.com/gophersaurus/gophersaurus/app/middleware"
)

// register takes a Router and registers route paths to controller methods.
func register(r *gf.Router) {

	// Register the HTTP GET pattern "/home" to the HomeController Index() method.
	r.Get("/home", c.Home.Index, middleware.Keys, gf.RespondErr)

	// Register the HTTP GET pattern "/work" to the WorkController Index() method.
	r.Get("/work", c.Work.Index, middleware.Keys, gf.RespondErr)

	// Register the pattern "/user" with all methods in the User Controller
	r.Resource("/user", "user_id", c.Users, middleware.SessionUser, middleware.SessionAdmin, middleware.Keys, gf.RespondErr)

	// Register the pattern "/user/{user_id}/address" with all methods in the Address Controller
	r.Resource("/user/{user_id}/address", "address_id", c.Addresses, middleware.SessionUser, middleware.SessionAdmin, middleware.Keys, gf.RespondErr)

	// Register
	r.Post("/session", c.Sessions.Store, middleware.Keys, gf.RespondErr)
	r.Get("/session", c.Sessions.Show, middleware.Keys, gf.RespondErr)

	// Serve static files.
	r.Static("/", "public/")
}
