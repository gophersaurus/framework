package models

import "git.target.com/gophersaurus/gophersaurus/vendor/git.target.com/gophersaurus/gf.v1"

type Creds struct {
	Username string `json:"username" val:"email"`
	Password string `json:"password" val:"password"`
}

var ValidatePassword func(v interface{}, param string) error

func init() {
	ValidatePassword = gf.NewPatternValidator("[a-zA-Z][a-zA-Z0-9]{7,19}$").Validate
	gf.Validator.SetValidationFunc("password", ValidatePassword)
}
