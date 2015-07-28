package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gophersaurus/framework/app/bootstrap"
	"github.com/gophersaurus/gf.v1/config"
	"github.com/gophersaurus/gf.v1/dba"
	"github.com/gophersaurus/gf.v1/router"
)

// Server describes a server application.
type Server struct {
	port   string
	static string
	keys   map[string][]string
	//TLS    config.TLS
}

// NewServer takes a config, databases, port, and keys and returns a new server.
func NewServer(
	port string,
	static string,
	// keys map[string][]string,
) Server {
	return Server{port: port, static: static /* , keys: keys , TLS: config.TLS*/}
}

// Bootstrap takes settings an returns a Server.
func Bootstrap() Server {
	bootstrap.Config()
	bootstrap.DB()
	return NewServer(config.GetString("port"), config.GetString("static"))
}

// Serve starts the application server.
func (s Server) Serve() {

	// defer closing db connections
	for _, db := range dba.All() {
		defer db.Close()
	}

	// mux
	m := router.NewMux()

	// keys
	if len(s.keys) > 0 {
		km := router.NewKeyMiddleware(s.keys)
		m.Middleware(km.Do)
	}

	// register dynamic routes
	register(m)

	// if a static directory path is provided, register it.
	if len(s.static) > 0 {
		m.Static("/public", s.static)
	} else {
		m.Static("/public", string(os.PathSeparator)+"public")
	}

	bootstrap.Docs(s.static, m.Endpoints())

	// serve and let the humans know...
	/*
		if len(s.TLS.Key) > 0 && len(s.TLS.Cert) > 0 {
			fmt.Println("\x1b[32;1m" + "Gophersaurus server listening with TLS on port :" + s.port + "\x1b[0m")
			log.Fatal(http.ListenAndServeTLS(":"+s.port, s.TLS.Cert, s.TLS.Key, m))
		} else {
	*/
	fmt.Println("\x1b[32;1m" + "Gophersaurus server listening on port :" + s.port + "\x1b[0m")
	log.Fatal(http.ListenAndServe(":"+s.port, m))
	//}
}
