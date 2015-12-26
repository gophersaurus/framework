package app

import (
	"github.com/gophersaurus/gf.v1/bootstrap"
	"github.com/gophersaurus/gf.v1/router"
)

// Serve bootstraps the web service.
func Serve() error {
	m := router.NewMux()
	return bootstrap.Server(m, register)
}
