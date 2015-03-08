package server

import (
	"fmt"
	"log"
	"net/http"

	"git.target.com/gophersaurus/gf.v1"
	"git.target.com/gophersaurus/gophersaurus/app/controllers"
	"git.target.com/gophersaurus/gophersaurus/app/middleware"
	"git.target.com/gophersaurus/gophersaurus/app/models"
	"git.target.com/gophersaurus/gophersaurus/bootstrap"
	"git.target.com/gophersaurus/gophersaurus/config"
)

// Server describes a server application.
type Server struct {
	port   string
	static string
	dba    *gf.DBA
	keys   map[string][]string
}

// NewServer takes a config, databases, port, and keys and returns a new server.
func NewServer(
	dba *gf.DBA,
	port string,
	static string,
	config config.Config,
	keys map[string][]string,
) Server {
	return Server{port: port, static: static, dba: dba, keys: keys}
}

// Bootstrap takes settings an returns a Server.
func Bootstrap(settings map[string]string) Server {

	// INITALIZE CONFIGURATION
	conf := bootstrap.Config(settings)

	// INITALIZE DATABASES
	dba := bootstrap.Databases(conf)

	// SERVER PORT
	port := conf.Port

	// STATIC FILE PATH
	static := settings["static"]

	// QUICK KEYS
	keys := conf.Keys

	// Initalize the database admin in models.
	models.Init(conf, dba)

	// Initalize the config object in controllers.
	controllers.Init(conf)

	// Initalize the config object in middleware.
	middleware.Init(conf)

	// SERVER
	return NewServer(dba, port, static, conf, keys)
}

// Serve starts the application server.
func (s Server) Serve() {

	fmt.Println("# STARTING SERVER")
	fmt.Println("	Defering closing databases connections...")

	// Defer the command to close Mongo db connections.
	for _, db := range s.dba.NoSQL {
		defer db.Close()
	}

	// Defer the command to close SQL db connections.
	for _, db := range s.dba.SQL {
		defer db.Close()
	}

	fmt.Println("	Creating a new router...")
	// Create a new router.
	r := gf.NewRouter(true)

	// If valid keys are provided, register them as gf.NewKeyMiddleware.
	if len(s.keys) > 0 {
		fmt.Println("	Attaching API keys middleware to router...")
		km := gf.NewKeyMiddleware(s.keys)
		r.Middleware(km.Do)
	}

	// register dynamic routes.
	fmt.Println("	Registering routes...")
	register(r)

	// If a static directory path is provided, register it.
	fmt.Println("	Setting static assets directory...")
	if len(s.static) > 0 {
		r.Static("/static", s.static)
	} else {
		r.Static("/static", "/public")
	}
	fmt.Print("\n")

	// let the humans know we are listening...
	fmt.Println("# SERVING")
	fmt.Println("	Server is listening on port " + s.port)
	fmt.Print("\n")
	log.Fatal(http.ListenAndServe(":"+s.port, r))
}
