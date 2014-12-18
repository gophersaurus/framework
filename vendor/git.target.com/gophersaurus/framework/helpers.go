package gophersauras

import (
	"errors"
	"log"
	"strconv"

	"git.target.com/gophersaurus/gophersaurus/vendor/gopkg.in/mgo.v2/bson"
)

func Check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// converts a string to a Mongo Id, returns an error if it cannot be converted
func StringToBsonId(id string) (bson.ObjectId, error) {
	if !bson.IsObjectIdHex(id) {
		return "", errors.New(InvalidId)
	}
	return bson.ObjectIdHex(id), nil
}

// converts a string to a int, returns an error if it cannot be converted
func StringToInt(id string) (int, error) {
	number, err := strconv.Atoi(id)
	if err != nil {
		return number, errors.New(InvalidParameter)
	}
	return number, nil
}
