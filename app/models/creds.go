package models

import gf "../../vendor/git.target.com/gophersaurus/framework"

type Creds struct {
	Username string `json:"username" val:"email"`
	Password string `json:"password" val:"password"`
}

func init() {
	val, err := gf.NewPatternValidator("[a-zA-Z][a-zA-Z0-9]{7,19}$")
	if err != nil {
		panic(err)
	}
	gf.Validator.SetValidationFunc("password", val.Validate)
}
