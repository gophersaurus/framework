package models

import "../../vendor/git.target.com/gophersaurus/framework"

type Creds struct {
	Username string `json:"username" validate:"email"`
	Password string `json:"password" validate:"password"`
}

var validateCreds = framework.Validate.BuildFunc(Creds{})

func (c *Creds) Validate() error {
	return validateCreds(*c)
}
