package controllers

import (
	"git.target.com/gophersaurus/gophersaurus/app/models"
	"git.target.com/gophersaurus/gophersaurus/app/repos"
	"git.target.com/gophersaurus/gophersaurus/app/validators"

	gf "git.target.com/gophersaurus/gophersaurus/vendor/git.target.com/gophersaurus/framework"
)

var User = &userController{}

type userController struct {
}

func (u *userController) Index(resp *gf.Response, req *gf.Request) {
	users, err := repos.FindAllUsers()
	if err != nil {
		// TODO -- how should we handle this error
		resp.RespondWithErr(err)
		return
	}
	resp.Body(users)
	resp.Respond()
}

func (u *userController) Store(resp *gf.Response, req *gf.Request) {
	user := models.NewUser()
	err := req.ReadBody(user)
	if err != nil {
		resp.RespondWithErr(err)
		return
	}
	err = gf.Validate(user)
	if err != nil {
		resp.RespondWithErr(err)
		return
	}
	err = user.Save()
	if err != nil {
		// TODO -- how should we handle this error
		resp.RespondWithErr(err)
		return
	}
	resp.Respond()
}

func (u *userController) Show(resp *gf.Response, req *gf.Request) {
	id, err := validators.ObjectId(req)
	if err != nil {
		resp.RespondWithErr(err)
		return
	}
	user := models.NewUser()
	err = user.Find("_id", id)
	if err != nil {
		// TODO -- how should we handle this error
		resp.RespondWithErr(err)
		return
	}
	resp.Body(user)
	resp.Respond()
}

func (u *userController) Update(resp *gf.Response, req *gf.Request) {
	id, err := validators.ObjectId(req)
	if err != nil {
		resp.RespondWithErr(err)
		return
	}
	user := models.NewUser()
	err = req.ReadBody(user)
	if err != nil {
		resp.RespondWithErr(err)
		return
	}
	user.Id = id
	err = gf.Validate(user)
	if err != nil {
		resp.RespondWithErr(err)
		return
	}
	err = user.Save()
	if err != nil {
		// TODO -- how should we handle this error
		resp.RespondWithErr(err)
		return
	}
	resp.Respond()
}

func (u *userController) Apply(resp *gf.Response, req *gf.Request) {
	id, err := validators.ObjectId(req)
	if err != nil {
		resp.RespondWithErr(err)
		return
	}
	user := models.NewUser()
	err = user.Find("_id", id)
	if err != nil {
		// TODO -- how should we handle this error
		resp.RespondWithErr(err)
		return
	}
	patch := Patch{}
	err = req.ReadBody(patch)
	if err != nil {
		resp.RespondWithErr(err)
		return
	}
	err = user.Apply(patch)
	if err != nil {
		resp.RespondWithErr(err)
		return
	}
	err = gf.Validate(user)
	if err != nil {
		resp.RespondWithErr(err)
		return
	}
	err = user.Save()
	if err != nil {
		// TODO -- how should we handle this error
		resp.RespondWithErr(err)
		return
	}
	resp.Respond()
}

func (u *userController) Destroy(resp *gf.Response, req *gf.Request) {
	id, err := validators.ObjectId(req)
	if err != nil {
		resp.RespondWithErr(err)
		return
	}
	user := models.NewUser()
	user.Id = id
	err = user.Delete()
	if err != nil {
		// TODO -- how should we handle this error
		resp.RespondWithErr(err)
		return
	}
	resp.Respond()
}
