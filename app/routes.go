package app

import (
	"git.target.com/gophersaurus/gf.v1"
	c "git.target.com/gophersaurus/gophersaurus/app/controllers"
	"git.target.com/gophersaurus/gophersaurus/app/middleware"
)

// register takes a Router and registers route paths to controller methods.
func register(keys gf.KeyMap, r *gf.Router) {

	keyHandler := middleware.NewKeyHandler(keys)

	sessionUser := middleware.NewSessionUserMiddleware("Session-Id", "user_id", "admin")

	// Register the HTTP GET pattern "/home" to the HomeController Index() method.
	r.Get("/home", c.Home.Index, keyHandler)

	// Register the HTTP GET pattern "/work" to the WorkController Index() method.
	r.Get("/work", c.Work.Index, keyHandler)

	// Register the pattern "/user" with all methods in the User Controller
	r.Resource("/user", "user_id", c.Users, keyHandler, sessionUser)

	// Register the pattern "/user/{user_id}/address" with all methods in the Address Controller
	r.Resource("/user/{user_id}/address", "address_id", c.Addresses, keyHandler, sessionUser)

	// Register
	r.Post("/session", c.Sessions.Store)
	r.Get("/session", c.Sessions.Show)

	// Serve static files.
	r.Static("/", "public/")
}
