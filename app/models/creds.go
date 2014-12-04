package models

import gf "../../vendor/git.target.com/gophersaurus/framework"

type Creds struct {
	Username string `json:"username" val:"email"`
	Password string `json:"password" val:"password"`
}

var ValidatePassword func(v interface{}, param string) error

func init() {
	ValidatePassword = gf.NewPatternValidator("[a-zA-Z][a-zA-Z0-9]{7,19}$").Validate
	gf.Validator.SetValidationFunc("password", ValidatePassword)
}
