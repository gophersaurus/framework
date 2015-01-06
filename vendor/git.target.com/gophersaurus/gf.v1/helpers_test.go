package gf

import (
	"testing"

	"git.target.com/gophersaurus/gophersaurus/vendor/git.target.com/gophersaurus/gf.v1/gophermocks"
	"git.target.com/gophersaurus/gophersaurus/vendor/git.target.com/gophersaurus/gf.v1/mockstar"
	"git.target.com/gophersaurus/gophersaurus/vendor/gopkg.in/mgo.v2/bson"
)

var tt *testing.T

func Test_StringToInt_Valid(t *testing.T) {
	mockstar.T = t

	value, err := StringToInt("17")

	mockstar.Expect(value).ToEqual(17)
	mockstar.Expect(err).ToBeNil()
}

func Test_StringToInt_Invalid(t *testing.T) {
	mockstar.T = t

	value, err := StringToInt("Steve")

	mockstar.Expect(value).ToEqual(0)
	mockstar.Expect(err.Error()).ToEqual(InvalidParameter)
}

func Test_StringToBsonId_Valid(t *testing.T) {
	mockstar.T = t

	idStr := "1234567890abcdef12345678" // this is a valid id
	value, err := StringToBsonId(idStr)

	mockstar.Expect(value).ToEqual(bson.ObjectIdHex(idStr))
	mockstar.Expect(err).ToBeNil()
}

func Test_StringToBsonId_Invalid(t *testing.T) {
	mockstar.T = t

	_, err := StringToBsonId("This is not a valid id")

	mockstar.Expect(err.Error()).ToEqual(InvalidId)
}

func Test_ObjectId_Exists(t *testing.T) {
	mockstar.T = t

	idStr := "1234567890abcdef12345678" // this is a valid id
	req := gophermocks.NewMockRequest()
	req.When("Var", "id").Return(idStr)
	value, err := ObjectId(req)

	mockstar.Expect(value).ToEqual(bson.ObjectIdHex(idStr))
	mockstar.Expect(err).ToBeNil()
}

func Test_ObjectId_DNE(t *testing.T) {
	mockstar.T = t

	req := gophermocks.NewMockRequest()
	req.When("Var", "id").Return("")
	_, err := ObjectId(req)

	mockstar.Expect(err.Error()).ToEqual(InvalidId)
}
