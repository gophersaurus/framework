package models

import "git.target.com/gophersaurus/framework"

type Creds struct {
	Username string `json:"username" val:"email"`
	Password string `json:"password" val:"password"`
}

var IsPasswordValid func(value string) bool

func init() {
	IsPasswordValid = gf.NewPatternValidator("[a-zA-Z][a-zA-Z0-9]{7,19}$").IsValid
	gf.ApplyLibValidator("password", IsPasswordValid)
}
