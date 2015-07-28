package bootstrap

import (
	"path/filepath"

	"github.com/gophersaurus/gf.v1/docs"
	"github.com/gophersaurus/gf.v1/router"
)

func Docs(static string, endpoints []router.Endpoint) {
	tmpl := filepath.Join(filepath.Dir(static), "app", "bootstrap", "templates", "endpoints.tmpl")
	html := filepath.Join(static, "docs", "api", "index.html")
	docs.APIendpoints(tmpl, html, endpoints)
}
