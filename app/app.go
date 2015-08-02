package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/gophersaurus/framework/app/bootstrap"
	"github.com/gophersaurus/framework/app/middleware"
	"github.com/gophersaurus/gf.v1/config"
	"github.com/gophersaurus/gf.v1/dba"
	"github.com/gophersaurus/gf.v1/router"
)

// Serve starts serving the application.
func Serve() {

	// bootstrap config and db
	bootstrap.Config()
	bootstrap.DB()

	// defer closing db connections
	for _, db := range dba.All() {
		defer db.Close()
	}

	// mux
	m := router.NewMux()

	// find config vars
	port := config.GetString("port")
	static := config.GetString("static")
	tls := config.GetStringMapString("tls")
	keys := config.GetStringMapStringSlice("keys")

	// key middleware
	if len(keys) > 0 {
		km := middleware.NewKeys(keys)
		m.Middleware(km.Do)
	}

	// register dynamic routes
	register(m)

	// if a static directory path is provided, register it
	if len(static) > 0 {
		m.Static("/public", static)
	} else {
		m.Static("/public", string(os.PathSeparator)+"public")
	}

	// bootstrap docs
	bootstrap.Docs(static, m.Endpoints())

	portStr := fmt.Sprintf(":%s", port)
	green := color.New(color.FgGreen).PrintfFunc()

	// serve and let the humans know...
	if tls["cert"] != "" && tls["key"] != "" {
		green("https listening and serving with TLS on port %s\n", port)
		log.Fatal(http.ListenAndServeTLS(portStr, tls["cert"], tls["key"], m))
	} else {
		green("http listening and serving on port %s\n", port)
		log.Fatal(http.ListenAndServe(portStr, m))
	}
}
