package app

import (
	c "github.com/gophersaurus/framework/app/controllers"
	"github.com/gophersaurus/gf.v1/router"
)

func register(r router.Router) {
	r.GET("/", c.Home.Index)
}
