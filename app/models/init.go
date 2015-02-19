package models

import "git.target.com/gophersaurus/gf.v1"

// The Database Administrator object.
var DBA gf.DBA

func Init(dba gf.DBA) {
	DBA = dba
}
