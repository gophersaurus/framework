package models

import "github.com/gophersaurus/gf.v1"

func init() {
	passValidator := gf.NewPatternValidator("[a-zA-Z][a-zA-Z0-9]{7,19}$")
	gf.Validator.SetValidationFunc("password", passValidator)
}

type Creds struct {
	Username string `json:"username" val:"email"`
	Password string `json:"password" val:"password"`
}
