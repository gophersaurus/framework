package resource

import (
	"errors"
	"strconv"
)

// MockModel represents a gf.Model.
type MockModel struct {
	ID    int
	data  string
	owner Modeler
}

// NewMockModel returns a *gf.MockModel.
func NewMockModel() *MockModel {
	return &MockModel{ID: 4037200794235010051, data: "mock model"}
}

// NewModel implements gf.Model.
func (m *MockModel) New() Modeler {
	return NewMockModel()
}

// PathID implements gf.Model.
func (m *MockModel) PathID() string {
	return "model_id"
}

// SetID implements gf.Model.
func (m *MockModel) SetID(id string) error {

	// try to convert the id to an integer
	i, err := strconv.Atoi(id)

	// error check
	if err != nil {
		return err
	}

	// set the id
	m.ID = i

	return nil
}

// FindAll implements gf.Model.
func (m *MockModel) FindAll() ([]Modeler, error) {

	// create a slice of gf.Model
	models := []Modeler{m.New(), m.New(), m.New()}

	return models, nil
}

// FindByID implements gf.Model.
func (m *MockModel) FindByID(id string) error {
	return nil
}

// FindAllByOwner implements gf.Model.
func (m *MockModel) FindAllByOwner(owner Modeler) ([]Modeler, error) {

	// create a slice of gf.Model
	models := []Modeler{m.New(), m.New(), m.New()}

	// range the models
	for _, model := range models {

		// attached each gf.Model to the owner
		model.BelongsTo(owner)
	}

	return models, nil
}

// BelongsTo implements gf.Model.
func (m *MockModel) BelongsTo(owner Modeler) error {

	// set the owner
	m.owner = owner

	return nil
}

// Save implements gf.Model.
func (m *MockModel) Save() error {
	return nil
}

// Delete implements gf.Model.
func (m *MockModel) Delete() error {
	return nil
}

// Validate implements gf.Model.
func (m *MockModel) Validate() error {
	return nil
}

// MockBaseModel represents a gf.Model.
type MockBaseModel struct {
	ID   int
	data string
}

// NewMockBaseModel returns a *gf.MockModel.
func NewMockBaseModel() *MockBaseModel {
	return &MockBaseModel{ID: 4037200794235010051, data: "mock model"}
}

// NewModel implements gf.Model.
func (m *MockBaseModel) New() Modeler {
	return NewMockBaseModel()
}

// PathID implements gf.Model.
func (m *MockBaseModel) PathID() string {
	return "base_id"
}

// SetID implements gf.Model.
func (m *MockBaseModel) SetID(id string) error {
	return nil
}

// FindAll implements gf.Model.
func (m *MockBaseModel) FindAll() ([]Modeler, error) {

	// create a slice of gf.Model
	models := []Modeler{m.New(), m.New(), m.New()}

	return models, nil
}

// FindByID implements gf.Model.
func (m *MockBaseModel) FindByID(id string) error {

	// try to convert the id to an integer
	i, err := strconv.Atoi(id)

	// error check
	if err != nil {
		return err
	}

	// set the id
	m.ID = i

	return nil
}

// FindAllByOwner implements gf.Model.
func (m *MockBaseModel) FindAllByOwner(owner Modeler) ([]Modeler, error) {
	return nil, errors.New("Base models usually do not have an owner.")
}

// BelongsTo implements gf.Model.
func (m *MockBaseModel) BelongsTo(owner Modeler) error {
	return errors.New("Base models usually do not have an owner.")
}

// Save implements gf.Model.
func (m *MockBaseModel) Save() error {
	return nil
}

// Delete implements gf.Model.
func (m *MockBaseModel) Delete() error {
	return nil
}

// Validate implements gf.Model.
func (m *MockBaseModel) Validate() error {
	return nil
}
