package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/gophersaurus/framework/app/bootstrap"
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

	// keys
	/*
		if len(keys) > 0 {
			km := middleware.NewKeys(keys)
			m.Middleware(km.Do)
		}
	*/

	// register dynamic routes
	register(m)

	// get port and static config vars
	port := config.GetString("port")
	static := config.GetString("static")

	// if a static directory path is provided, register it
	if len(static) > 0 {
		m.Static("/public", static)
	} else {
		m.Static("/public", string(os.PathSeparator)+"public")
	}

	// bootstrap docs
	bootstrap.Docs(static, m.Endpoints())

	// serve and let the humans know...
	/*
		if len(s.TLS.Key) > 0 && len(s.TLS.Cert) > 0 {
			fmt.Println("\x1b[32;1m" + "gophersaurus server listening with TLS on port :" + s.port + "\x1b[0m")
			log.Fatal(http.ListenAndServeTLS(":"+s.port, s.TLS.Cert, s.TLS.Key, m))
		} else {
	*/

	portStr := fmt.Sprintf(":%s", port)
	green := color.New(color.FgGreen).PrintfFunc()
	green("gophersaurus server listening on %s\n", portStr)
	log.Fatal(http.ListenAndServe(portStr, m))
	//}
}
