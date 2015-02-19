package server

import (
	"log"
	"net/http"

	"git.target.com/gophersaurus/gf.v1"
	"git.target.com/gophersaurus/gophersaurus/app/controllers"
	"git.target.com/gophersaurus/gophersaurus/app/models"
	"git.target.com/gophersaurus/gophersaurus/bootstrap"
	"git.target.com/gophersaurus/gophersaurus/config"
)

// Server describes a server application.
type Server struct {
	port   string
	static string
	keys   map[string][]string
}

// NewServer takes a config, databases, port, and keys and returns a new server.
func NewServer(
	dba gf.DBA,
	port string,
	static string,
	config config.Config,
	keys map[string][]string,
) Server {

	// Initalize the database admin in models.
	models.Init(dba)

	// Initalize the config object in controllers.
	controllers.Init(config)

	// Return a new Server.
	return Server{port: port, static: static, keys: keys}
}

// Bootstrap takes settings an returns a Server.
func Bootstrap(settings map[string]string) Server {

	// INITALIZE CONFIGURATION
	c := bootstrap.Config(settings)

	// INITALIZE DATABASES
	dba := bootstrap.Databases(c)

	// SERVER PORT
	p := c.Port

	// STATIC FILE PATH
	s := settings["static"]

	// QUICK KEYS
	k := c.Keys

	// SERVER
	return NewServer(dba, p, s, c, k)
}

// Serve starts the application server.
func (s Server) Serve() {

	// Defer the command to close db connections after HTTP execution completes.
	for _, db := range models.DBA {
		defer db.Close()
	}

	// Create a new router.
	r := gf.NewRouter()

	// If valid keys are provided, register them as gf.NewKeyMiddleware.
	if len(s.keys) > 0 {
		r.Middleware(gf.NewKeyMiddleware(s.keys))
	}

	// If a static directory path is provided, register it.
	if len(s.static) > 0 {
		r.Static("/", s.static)
	} else {
		r.Static("/", "/public")
	}

	// register dynamic routes.
	register(r)

	// let the humans know we are listening...
	log.Println("listening on port " + s.port)
	log.Fatal(http.ListenAndServe(":"+s.port, r))
}
