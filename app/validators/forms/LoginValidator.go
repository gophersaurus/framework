package forms

import (
	f "../../../vendor/git.target.com/gospot/framework"
	"../../models"
)

func Login(req *f.Request) error {
	creds := models.Creds{}
	err := req.ReadBody(&creds)
	if err != nil {
		return err
	}
	return (&creds).Validate()
}
