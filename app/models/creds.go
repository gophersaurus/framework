package models

import "github.com/gophersaurus/gf.v1"

// Creds describes user credentials.
type Creds struct {
	Username string `json:"username" val:"email"`
	Password string `json:"password" val:"password"`
}

func init() {
	passRegex := gf.NewPatternValidator("[a-zA-Z][a-zA-Z0-9]{7,19}$")
	gf.Validator.SetValidationFunc("password", passRegex)
}
