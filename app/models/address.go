package models

import (
	"errors"

	"github.com/gophersaurus/gf.v1"
	"gopkg.in/mgo.v2/bson"
)

type Address struct {
	user   *User
	Id     string `json:"id" bson:"id"`
	Street string `json:"street" bson:"street"`
	City   string `json:"city" bson:"city"`
	State  string `json:"state" bson:"state"`
	Zip    string `json:"zip" bson:"zip"`
}

func NewAddress() *Address {
	return &Address{Id: bson.NewObjectId().Hex()}
}

func (a *Address) NewModel() gf.Model {
	return &Address{Id: bson.NewObjectId().Hex()}
}

func (a *Address) PathID() string {
	return "address_id"
}

func (a *Address) SetID(id string) error {
	a.Id = id
	return nil
}

func (a *Address) BelongsTo(owner gf.Model) error {
	if user, ok := owner.(*User); ok {
		a.user = user
		return nil
	}
	return errors.New("invalid parent type")
}

func (a *Address) FindByID(id string) error {
	index, err := a.find(id)
	if err != nil {
		return err
	}
	if index < 0 {
		return errors.New("not found")
	}
	user := a.user
	*a = a.user.Addresses[index]
	a.user = user
	return nil
}

func (a *Address) FindAll() ([]gf.Model, error) {

	// Create a slice of addresses to work with.
	var addresses []Address

	// Search the databases for addresses.
	if err := dba.MGO("test").C("testAddresses").Find(bson.M{}).All(&addresses); err != nil {
		return nil, err
	}

	// Unfortunately a []struct that implements []gf.Model is not compatible.
	// This is because of how interface memory is mapped in golang. It sucks.
	// More here: https://groups.google.com/forum/#!topic/golang-nuts/Il-tO1xtAyE
	models := make([]gf.Model, len(addresses))
	for i, v := range addresses {
		models[i] = gf.Model(&v)
	}

	return models, nil
}

func (a *Address) Save() error {
	index, err := a.find(a.Id)
	if err != nil {
		return err
	}
	if index < 0 {
		a.user.Addresses = append(a.user.Addresses, *a)
	} else {
		a.user.Addresses[index] = *a
	}
	a.user.Save()
	return nil
}

func (a *Address) Delete() error {
	index, err := a.find(a.Id)
	if err != nil {
		return err
	}
	if index < 0 {
		return errors.New("not found")
	}
	addresses := []Address{}
	if index > 0 {
		addresses = append(addresses, a.user.Addresses[:index]...)
	}
	if index+1 < len(a.user.Addresses) {
		addresses = append(addresses, a.user.Addresses[index+1:]...)
	}
	a.user.Addresses = addresses
	a.user.Save()
	return nil
}

func (a *Address) FindAllByOwner(owner gf.Model) ([]gf.Model, error) {

	// Create a new array of models.
	models := []gf.Model{}

	if user, ok := owner.(*User); ok {

		// Range through addresses.
		for _, address := range user.Addresses {
			temp := address // needed for pointer
			models = append(models, &temp)
		}

		return models, nil
	}
	return nil, errors.New("invalid parent type")
}

func (a *Address) Validate() error {
	return gf.Validate(a)
}

func (a *Address) find(id string) (int, error) {
	if a.user == nil {
		return 0, errors.New("parent not set")
	}
	for index, address := range a.user.Addresses {
		if address.Id == id {
			return index, nil
		}
	}
	return -1, nil
}
