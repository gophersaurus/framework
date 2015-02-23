package models

import (
	"git.target.com/gophersaurus/gf.v1"
	"git.target.com/gophersaurus/gophersaurus/config"
)

var dba *gf.DBA
var conf config.Config

func Init(c config.Config, d *gf.DBA) {
	dba = d
	conf = c
}
