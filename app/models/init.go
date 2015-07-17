package models

import (
	"github.com/gophersaurus/framework/config"
	"github.com/gophersaurus/gf.v1/database"
)

var dba *database.DBA
var conf config.Config

func Init(c config.Config, d *database.DBA) {
	dba = d
	conf = c
}
