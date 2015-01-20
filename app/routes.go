package app

import (
	"git.target.com/gophersaurus/gf.v1"
	c "git.target.com/gophersaurus/gophersaurus/app/controllers"
	m "git.target.com/gophersaurus/gophersaurus/app/middleware"
)

// register takes a Router and registers route paths to controller methods.
func register(r *gf.Router) {

	// Register the HTTP GET pattern "/home" to the HomeController Index() method.
	r.Get("/home", c.Home.Index, m.Keys, gf.RespondErr)

	// Register the HTTP GET pattern "/work" to the WorkController Index() method.
	r.Get("/work", c.Work.Index, m.Keys, gf.RespondErr)

	// Register the pattern "/user" with all methods in the User Controller
	r.Resource("/user", "user_id", c.Users, m.Keys, m.SessionUserAdmin, gf.RespondErr)

	// Register the pattern "/user/{user_id}/address" with all methods in the Address Controller
	r.Resource("/user/{user_id}/address", "address_id", c.Addresses, m.Keys, m.SessionUserAdmin, gf.RespondErr)

	// Register
	r.Post("/session", c.Sessions.Store, m.Keys, gf.RespondErr)
	r.Get("/session", c.Sessions.Show, m.Keys, gf.RespondErr)

	// Serve static files.
	r.Static("/", "public/")
}
