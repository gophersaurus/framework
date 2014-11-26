package gophersauras

import (
	"errors"
	"log"
	"strconv"

	"../../../gopkg.in/mgo.v2/bson"
)

func Check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// converts a string to a Mongo Id, returns an error if it cannot be converted
func ToBsonId(id string) (bson.ObjectId, error) {
	if !bson.IsObjectIdHex(id) {
		return "", errors.New(InvalidId)
	}
	return bson.ObjectIdHex(id), nil
}

// converts a tcin string to a int, returns an error if it cannot be converted
func ToInt(id string) (int, error) {
	number, err := strconv.Atoi(id)
	if err != nil {
		return number, errors.New(InvalidParameter)
	}
	return number, nil
}
