package server

import (
	"fmt"
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
	config *config.Config,
	keys map[string][]string,
) Server {

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

	// Initalize the database admin in models.
	models.Init(dba)

	// Initalize the config object in controllers.
	controllers.Init(config)

	// SERVER
	return NewServer(dba, p, s, c, k)
}

// Serve starts the application server.
func (s Server) Serve() {

	fmt.Println("# STARTING SERVER")
	fmt.Println("	Defering closing databases connections...")

	// Defer the command to close Mongo db connections.
	for _, db := range models.DBA.NoSQL {
		defer db.Close()
	}

	// Defer the command to close SQL db connections.
	for _, db := range models.DBA.SQL {
		defer db.Close()
	}

	fmt.Println("	Creating a new router...")
	// Create a new router.
	r := gf.NewRouter()

	// If valid keys are provided, register them as gf.NewKeyMiddleware.
	if len(s.keys) > 0 {
		fmt.Println("	Attaching API keys middleware to router...")
		r.Middleware(gf.NewKeyMiddleware(s.keys))
	}

	// register dynamic routes.
	fmt.Println("	Registering routes...")
	register(r)

	// If a static directory path is provided, register it.
	fmt.Println("	Setting static assets directory...")
	if len(s.static) > 0 {
		r.Static("/", s.static)
	} else {
		r.Static("/", "/public")
	}

	fmt.Print("\n")

	// let the humans know we are listening...
	fmt.Println("# SERVING")
	fmt.Println("	Server is listening on port " + s.port)
	fmt.Print("\n")
	log.Fatal(http.ListenAndServe(":"+s.port, r))
}
