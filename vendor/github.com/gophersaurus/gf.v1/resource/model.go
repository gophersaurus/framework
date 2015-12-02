package resource

// Modeler represents a data model used by resource controllers to automate
// RESTful CRUD operations.
type Modeler interface {
	New() Modeler
	PathID() string
	SetID(id string) error
	FindAll() ([]Modeler, error)
	FindByID(id string) error
	FindAllByOwner(owner Modeler) ([]Modeler, error)
	BelongsTo(owner Modeler) error
	Save() error
	Delete() error
	Validate() error
}
