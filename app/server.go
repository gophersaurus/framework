package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gophersaurus/framework/app/controllers"
	"github.com/gophersaurus/framework/app/middleware"
	"github.com/gophersaurus/framework/app/models"
	"github.com/gophersaurus/framework/bootstrap"
	"github.com/gophersaurus/framework/config"
	"github.com/gophersaurus/gf.v1"
)

// Server describes a server application.
type Server struct {
	port   string
	static string
	dba    *gf.DBA
	keys   map[string][]string
	TLS    gf.TLS
}

// NewServer takes a config, databases, port, and keys and returns a new server.
func NewServer(
	dba *gf.DBA,
	port string,
	static string,
	config config.Config,
	keys map[string][]string,
) Server {
	return Server{port: port, static: static, dba: dba, keys: keys, TLS: config.TLS}
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

	// Defer the command to close Mongo db connections.
	for _, db := range s.dba.NoSQL {
		defer db.Close()
	}

	// Defer the command to close SQL db connections.
	for _, db := range s.dba.SQL {
		defer db.Close()
	}

	// Create a new router.
	r := gf.NewRouter(true)

	// If valid keys are provided, register them as gf.NewKeyMiddleware.
	if len(s.keys) > 0 {
		km := gf.NewKeyMiddleware(s.keys)
		r.Middleware(km.Do)
	}

	// register dynamic routes.
	register(r)

	// If a public static directory path is provided, register it.
	if len(s.static) > 0 {
		r.Static("/public", s.static)
	} else {
		r.Static("/public", string(os.PathSeparator)+"public")
	}

	// auto generate API documentation
	r.GenAPIDoc(filepath.Join(filepath.Dir(s.static), "bootstrap", "apidoc.tmpl"), "public/docs/api/index.html")

	// serve and let the humans know we are serving...
	if len(s.TLS.Key) > 0 && len(s.TLS.Cert) > 0 {
		fmt.Println("\x1b[32;1m" + "Gophersaurus server listening with TLS on port :" + s.port + "\x1b[0m")
		log.Fatal(http.ListenAndServeTLS(":"+s.port, s.TLS.Cert, s.TLS.Key, r))
	} else {
		fmt.Println("\x1b[32;1m" + "Gophersaurus server listening on port :" + s.port + "\x1b[0m")
		log.Fatal(http.ListenAndServe(":"+s.port, r))
	}
}
