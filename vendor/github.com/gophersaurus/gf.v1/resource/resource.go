// Package resource provides methods to automate CRUD via resource controllers.
package resource

import (
	"errors"
	"log"

	"github.com/gophersaurus/gf.v1/http"
)

// Resourcer represents a resource controller and its action methods.
type Resourcer interface {

	// action methods
	Index(resp http.ResponseWriter, req *http.Request)
	Show(resp http.ResponseWriter, req *http.Request)
	Store(resp http.ResponseWriter, req *http.Request)
	Apply(resp http.ResponseWriter, req *http.Request)
	Update(resp http.ResponseWriter, req *http.Request)
	Destroy(resp http.ResponseWriter, req *http.Request)
}

// Resource describes a resource controller.
type Resource struct {
	model Modeler
	ID    func(req *http.Request) (string, error)
}

// New takes a model and returns a new Resource.
func New(m Modeler,

	// Optional way to get an ID from request.
	// This is useful for '/user' routes based on session identification.
	optID ...func(req *http.Request) (string, error),

) *Resource {

	// new resource
	r := &Resource{model: m,
		ID: func(req *http.Request) (string, error) {
			return PathID(req, m.PathID())
		},
	}

	// set given optional id
	if len(optID) > 0 {
		r.ID = optID[0]
	}

	return r
}

// Index is a GET request for returning a list of items.
func (r *Resource) Index(resp http.ResponseWriter, req *http.Request) {

	// find list of items
	items, err := r.model.FindAll()

	// log the error, don't return it
	if err != nil {
		log.Fatal(err)
	}

	// write item list
	resp.WriteFormatList(req, items)
}

// Show is a GET request for displaying a single item.
func (r *Resource) Show(resp http.ResponseWriter, req *http.Request) {

	// get the resource id
	id, err := r.ID(req)
	if err != nil {
		resp.WriteErrs(req, err)
		return
	}

	item := r.model.New()

	// find the item by id
	if err := item.FindByID(id); err != nil {
		resp.WriteErrs(req, err)
		return
	}

	resp.WriteFormat(req, item)
}

// Store is a POST request for creating a new item.
func (r *Resource) Store(resp http.ResponseWriter, req *http.Request) {

	item := r.model.New()

	// read POST body for item data
	if err := req.UnmarshalJSONBody(item); err != nil {
		resp.WriteErrs(req, err)
		return
	}

	if err := item.Validate(); err != nil {
		resp.WriteErrs(req, err)
		return
	}

	if err := item.Save(); err != nil {
		log.Fatal(err)
	}

	resp.Status(http.StatusCreated)
	resp.WriteFormat(req, item)
}

// Apply is a PATCH request for updating a single item.
func (r *Resource) Apply(resp http.ResponseWriter, req *http.Request) {

	// get the id
	id, err := r.ID(req)
	if err != nil {
		resp.WriteErrs(req, err)
		return
	}

	item := r.model.New()

	if err := item.FindByID(id); err != nil {
		resp.WriteErrs(req, err)
		return
	}

	if err := req.UnmarshalJSONBody(item); err != nil {
		resp.WriteErrs(req, err)
		return
	}

	if err := item.Validate(); err != nil {
		resp.WriteErrs(req, err)
		return
	}

	if err := item.Save(); err != nil {
		resp.WriteErrs(req, err)
		return
	}

	resp.WriteFormat(req, item)
}

// Update is a PUT request for replacing a single item.
func (r *Resource) Update(resp http.ResponseWriter, req *http.Request) {

	// get resource id
	id, err := r.ID(req)
	if err != nil {
		resp.WriteErrs(req, err)
		return
	}

	item := r.model.New()

	// read PUT body for item data
	if err := req.UnmarshalJSONBody(item); err != nil {
		resp.WriteErrs(req, err)
		return
	}

	if err := item.SetID(id); err != nil {
		resp.WriteErrs(req, err)
		return
	}

	if err := item.Validate(); err != nil {
		resp.WriteErrs(req, err)
		return
	}

	if err := item.Save(); err != nil {
		resp.WriteErrs(req, err)
		return
	}

	resp.WriteFormat(req, item)
}

// Destroy is a DELETE request for deleting a single item.
func (r *Resource) Destroy(resp http.ResponseWriter, req *http.Request) {

	// get resource id
	id, err := r.ID(req)
	if err != nil {
		resp.WriteErrs(req, err)
		return
	}

	item := r.model.New()

	if err := item.SetID(id); err != nil {
		resp.WriteErrs(req, err)
		return
	}

	if err := item.Delete(); err != nil {
		resp.WriteErrs(req, err)
		return
	}

	resp.Status(http.StatusNoContent)
	resp.WriteFormat(req, "")
}

// PathID finds an id in a request url.
func PathID(req *http.Request, pathID string) (string, error) {
	if id := req.Param(pathID); len(id) > 0 {
		return id, nil
	}
	return "", errors.New(http.InvalidID)
}
