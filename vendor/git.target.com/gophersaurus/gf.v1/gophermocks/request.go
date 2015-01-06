package gophermocks

import (
	"fmt"
	"net/http"

	"git.target.com/gophersaurus/gophersaurus/vendor/git.target.com/gophersaurus/gf.v1/mockstar"
	"git.target.com/gophersaurus/gophersaurus/vendor/gopkg.in/mgo.v2/bson"
)

type mockRequest struct {
	*mockstar.Mock
}

func NewMockRequest() *mockRequest {
	return &mockRequest{mockstar.NewMock()}
}

func (m *mockRequest) Request() *http.Request {
	args := m.Mock.Called("Request")
	return m.castRequest(args.Get(0))
}

func (m *mockRequest) castRequest(obj interface{}) *http.Request {
	if obj == nil {
		return nil
	}
	out, ok := obj.(*http.Request)
	if !ok {
		mockstar.Fatal("cannot cast value to http.Request: " + fmt.Sprintf("%v", obj))
	}
	return out
}

func (m *mockRequest) Var(name string) string {
	args := m.Mock.Called("Var", name)
	return args.String(0)
}

func (m *mockRequest) Query(name string) []string {
	args := m.Mock.Called("Query", name)
	return args.Strings(0)
}

func (m *mockRequest) SessionId() bson.ObjectId {
	args := m.Mock.Called("SessionId")
	return m.castBsonId(args.Get(0))
}

func (m *mockRequest) castBsonId(obj interface{}) bson.ObjectId {
	if obj == nil {
		return ""
	}
	out, ok := obj.(bson.ObjectId)
	if !ok {
		mockstar.Fatal("cannot cast value to bson.ObjectId: " + fmt.Sprintf("%v", obj))
	}
	return out
}

func (m *mockRequest) Body() string {
	args := m.Mock.Called("Body")
	return args.String(0)
}

func (m *mockRequest) HasSession() bool {
	args := m.Mock.Called("HasSession")
	return args.Bool(0)
}

func (m *mockRequest) HasBody() bool {
	args := m.Mock.Called("HasBody")
	return args.Bool(0)
}

func (m *mockRequest) ReadBody(obj interface{}) error {
	args := m.Mock.Called("ReadBody", obj)
	return args.Err(0)
}
