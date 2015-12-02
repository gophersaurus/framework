package app

import (
	"github.com/gophersaurus/framework/app/middleware"
	"github.com/gophersaurus/gf.v1/bootstrap"
	"github.com/gophersaurus/gf.v1/config"
	"github.com/gophersaurus/gf.v1/router"
)

// Serve starts serving the web application.
func Serve() error {
	if err := bootstrap.Config(); err != nil {
		return err
	}

	m := router.NewMux()
	keys := config.GetStringMapStringSlice("keys")

	if len(keys) > 0 {
		km := middleware.NewKeys(keys)
		m.Middleware(km.Do)
	}

	return bootstrap.Server(m, register)
}
