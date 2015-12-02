package bootstrap

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/gophersaurus/gf.v1/config"
	"github.com/gophersaurus/gf.v1/dba"
	"github.com/gophersaurus/gf.v1/docs"
	"github.com/gophersaurus/gf.v1/router"
)

// Server takes a register function and bootstraps a server.
func Server(r router.Router, register func(r router.Router)) error {

	// bootstrap database admin
	if err := DBA(); err != nil {
		return err
	}

	// defer closing dba connections
	for _, db := range dba.All() {
		defer db.Close()
	}

	register(r)

	port := config.GetString("port")
	static := config.GetString("static")
	tls := config.GetStringMapString("tls")

	// if a static directory path is provided, register it
	if len(static) > 0 {
		r.Static("/public", static)
	} else {
		r.Static("/public", string(os.PathSeparator)+"public")
	}

	// generate docs
	if err := Docs(static, router.Endpoints()); err != nil {
		return err
	}

	// prep port and green output
	portStr := fmt.Sprintf(":%s", port)
	green := color.New(color.FgGreen).PrintfFunc()

	// let the humans know we are serving...
	if tls["cert"] != "" && tls["key"] != "" {
		green("https listening and serving with TLS on port %s\n", port)
		return http.ListenAndServeTLS(portStr, tls["cert"], tls["key"], r)
	}

	green("http listening and serving on port %s\n", port)
	return http.ListenAndServe(portStr, r)
}

// Docs renders all the endpoint docs for the API application service.
func Docs(static string, endpoints []router.Endpoint) error {
	tmpl := filepath.Join(filepath.Dir(static), "app", "templates", "endpoints.tmpl")
	html := filepath.Join(static, "docs", "api", "index.html")
	if err := docs.Endpoints(tmpl, html, endpoints); err != nil {
		return err
	}
	return nil
}
