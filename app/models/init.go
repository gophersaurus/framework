package models

import (
	"github.com/gophersaurus/gf.v1"
	"github.com/gophersaurus/framework/config"
)

var dba *gf.DBA
var conf config.Config

func Init(c config.Config, d *gf.DBA) {
	dba = d
	conf = c
}
