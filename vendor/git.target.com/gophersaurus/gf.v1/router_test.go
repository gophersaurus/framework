package gf

import (
	"net/http"
	"strings"
	"testing"

	"git.target.com/gophersaurus/gophersaurus/vendor/git.target.com/gophersaurus/gf.v1/gophermocks"
	"git.target.com/gophersaurus/gophersaurus/vendor/git.target.com/gophersaurus/gf.v1/mockstar"
)

func Test_Router_Get(t *testing.T) {
	mockstar.T = t

	ctrl := newMockController()
	ctrl.When("Index", mockstar.Any, mockstar.Any).Return()

	writer := gophermocks.NewMockResponseWriter()
	writer.When("WriteHeader", http.StatusOK).Return()

	key := "This_is_a_key"
	path := "/mock"

	router := NewRouter(map[Key]KeyConfig{
		Key(key): KeyConfig{true, []string{}},
	}, false)

	router.Get(path, ctrl.Index)

	domain := "http://1.2.3.4:5678"
	req, err := http.NewRequest("GET", domain+path, strings.NewReader(""))
	mockstar.Expect(err).ToBeNil()
	req.Header.Set("API-Key", key)

	router.ServeHTTP(writer, req)

	call := ctrl.GetCall("Index", mockstar.Any, mockstar.Any)
	count := call.GetCallCount()
	mockstar.Expect(count).ToEqual(1)
	args := call.GetCallParams(0)
	resp, ok := args.Get(0).(Response)
	mockstar.Expect(ok).ToBeTrue()

	mockstar.Expect(writer.HasCalled("WriteHeader", http.StatusOK).Times(0)).ToBeTrue()
	resp.Respond()
	mockstar.Expect(writer.HasCalled("WriteHeader", http.StatusOK).Once()).ToBeTrue()

	reqObj, ok := args.Get(1).(Request)
	mockstar.Expect(ok).ToBeTrue()
	mockstar.Expect(reqObj.Request()).ToEqual(req)
}

func Test_Router_DeadEnd(t *testing.T) {
	mockstar.T = t

	writer := gophermocks.NewMockResponseWriter()
	responseHeader := http.Header{}
	writer.When("Header").Return(responseHeader)
	writer.When("WriteHeader", http.StatusNotFound).Return()
	expectedErr := []byte("404 page not found\n")
	writer.When("Write", expectedErr).Return()

	router := NewRouter(map[Key]KeyConfig{}, false)

	domain := "http://1.2.3.4:5678/mock"
	req, err := http.NewRequest("GET", domain, strings.NewReader(""))
	mockstar.Expect(err).ToBeNil()

	router.ServeHTTP(writer, req)

	mockstar.Expect(writer.HasCalled("WriteHeader", http.StatusNotFound).Once()).ToBeTrue()
	mockstar.Expect(writer.HasCalled("Write", expectedErr).Once()).ToBeTrue()

	mockstar.Expect(len(responseHeader)).ToEqual(1)
	mockstar.Expect(responseHeader.Get("Content-Type")).ToEqual("text/plain; charset=utf-8")
}

func Test_Router_Get_BadKey(t *testing.T) {
	mockstar.T = t

	ctrl := newMockController()
	ctrl.When("Index", mockstar.Any, mockstar.Any).Return()

	writer := gophermocks.NewMockResponseWriter()
	responseHeader := http.Header{}
	writer.When("Header").Return(responseHeader)
	writer.When("WriteHeader", http.StatusForbidden).Return()
	expectedErr := []byte("{\"error\":\"invalid permissions\"}")
	writer.When("Write", expectedErr).Return()

	key := "This_is_a_key"
	path := "/mock"

	router := NewRouter(map[Key]KeyConfig{
		Key(key): KeyConfig{true, []string{}},
	}, false)

	router.Get(path, ctrl.Index)

	domain := "http://1.2.3.4:5678"
	req, err := http.NewRequest("GET", domain+path, strings.NewReader(""))
	mockstar.Expect(err).ToBeNil()

	router.ServeHTTP(writer, req)

	mockstar.Expect(ctrl.HasCalled("Index", mockstar.Any, mockstar.Any).Times(0)).ToBeTrue()

	mockstar.Expect(writer.HasCalled("WriteHeader", http.StatusForbidden).Once()).ToBeTrue()
	mockstar.Expect(writer.HasCalled("Write", expectedErr).Once()).ToBeTrue()

	mockstar.Expect(len(responseHeader)).ToEqual(1)
	mockstar.Expect(responseHeader.Get("Content-Type")).ToEqual("application/json; charset=UTF-8")
}

func Test_Router_Post(t *testing.T) {
	mockstar.T = t

	ctrl := newMockController()
	ctrl.When("Store", mockstar.Any, mockstar.Any).Return()

	writer := gophermocks.NewMockResponseWriter()
	writer.When("WriteHeader", http.StatusOK).Return()

	key := "This_is_a_key"
	path := "/mock"

	router := NewRouter(map[Key]KeyConfig{
		Key(key): KeyConfig{true, []string{}},
	}, false)

	router.Post(path, ctrl.Store)

	domain := "http://1.2.3.4:5678"
	req, err := http.NewRequest("POST", domain+path, strings.NewReader(""))
	mockstar.Expect(err).ToBeNil()
	req.Header.Set("API-Key", key)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(writer, req)

	call := ctrl.GetCall("Store", mockstar.Any, mockstar.Any)
	count := call.GetCallCount()
	mockstar.Expect(count).ToEqual(1)
	args := call.GetCallParams(0)
	resp, ok := args.Get(0).(Response)
	mockstar.Expect(ok).ToBeTrue()

	mockstar.Expect(writer.HasCalled("WriteHeader", http.StatusOK).Times(0)).ToBeTrue()
	resp.Respond()
	mockstar.Expect(writer.HasCalled("WriteHeader", http.StatusOK).Once()).ToBeTrue()

	reqObj, ok := args.Get(1).(Request)
	mockstar.Expect(ok).ToBeTrue()
	mockstar.Expect(reqObj.Request()).ToEqual(req)
}

func Test_Router_Post_NoContentType(t *testing.T) {
	mockstar.T = t

	ctrl := newMockController()
	ctrl.When("Store", mockstar.Any, mockstar.Any).Return()

	writer := gophermocks.NewMockResponseWriter()
	responseHeader := http.Header{}
	writer.When("Header").Return(responseHeader)
	writer.When("WriteHeader", http.StatusNotFound).Return()
	expectedErr := []byte("404 page not found\n")
	writer.When("Write", expectedErr).Return()

	key := "This_is_a_key"
	path := "/mock"

	router := NewRouter(map[Key]KeyConfig{
		Key(key): KeyConfig{true, []string{}},
	}, false)

	router.Post(path, ctrl.Store)

	domain := "http://1.2.3.4:5678"
	req, err := http.NewRequest("POST", domain+path, strings.NewReader(""))
	mockstar.Expect(err).ToBeNil()
	req.Header.Set("API-Key", key)

	router.ServeHTTP(writer, req)

	mockstar.Expect(ctrl.HasCalled("Store", mockstar.Any, mockstar.Any).Times(0)).ToBeTrue()

	mockstar.Expect(writer.HasCalled("WriteHeader", http.StatusNotFound).Once()).ToBeTrue()
	mockstar.Expect(writer.HasCalled("Write", expectedErr).Once()).ToBeTrue()

	mockstar.Expect(len(responseHeader)).ToEqual(1)
	mockstar.Expect(responseHeader.Get("Content-Type")).ToEqual("text/plain; charset=utf-8")

}

func Test_Router_Resource(t *testing.T) {
	mockstar.T = t

	ctrl := newMockController()
	ctrl.When("Index", mockstar.Any, mockstar.Any).Return()
	ctrl.When("Store", mockstar.Any, mockstar.Any).Return()
	ctrl.When("Show", mockstar.Any, mockstar.Any).Return()
	ctrl.When("Destroy", mockstar.Any, mockstar.Any).Return()

	writer := gophermocks.NewMockResponseWriter()
	writer.When("WriteHeader", http.StatusOK).Return()

	key := "This_is_a_key"
	path := "/mock"
	id := "This_is_the_id"

	router := NewRouter(map[Key]KeyConfig{
		Key(key): KeyConfig{true, []string{}},
	}, false)

	router.Resource(path, ctrl)

	domain := "http://1.2.3.4:5678"

	// check ctrl.Index
	req, err := http.NewRequest("GET", domain+path, strings.NewReader(""))
	mockstar.Expect(err).ToBeNil()
	req.Header.Set("API-Key", key)

	router.ServeHTTP(writer, req)

	call := ctrl.GetCall("Index", mockstar.Any, mockstar.Any)
	count := call.GetCallCount()
	mockstar.Expect(count).ToEqual(1)
	args := call.GetCallParams(0)
	resp, ok := args.Get(0).(Response)
	mockstar.Expect(ok).ToBeTrue()

	mockstar.Expect(writer.HasCalled("WriteHeader", http.StatusOK).Times(0)).ToBeTrue()
	resp.Respond()
	mockstar.Expect(writer.HasCalled("WriteHeader", http.StatusOK).Once()).ToBeTrue()

	reqObj, ok := args.Get(1).(Request)
	mockstar.Expect(ok).ToBeTrue()
	mockstar.Expect(reqObj.Request()).ToEqual(req)
	mockstar.Expect(reqObj.Var("id")).ToEqual("")

	// check ctrl.Store
	req, err = http.NewRequest("POST", domain+path, strings.NewReader(""))
	mockstar.Expect(err).ToBeNil()
	req.Header.Set("API-Key", key)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(writer, req)

	call = ctrl.GetCall("Store", mockstar.Any, mockstar.Any)
	count = call.GetCallCount()
	mockstar.Expect(count).ToEqual(1)
	args = call.GetCallParams(0)
	resp, ok = args.Get(0).(Response)
	mockstar.Expect(ok).ToBeTrue()

	mockstar.Expect(writer.HasCalled("WriteHeader", http.StatusOK).Once()).ToBeTrue()
	resp.Respond()
	mockstar.Expect(writer.HasCalled("WriteHeader", http.StatusOK).Twice()).ToBeTrue()

	reqObj, ok = args.Get(1).(Request)
	mockstar.Expect(ok).ToBeTrue()
	mockstar.Expect(reqObj.Request()).ToEqual(req)
	mockstar.Expect(reqObj.Var("id")).ToEqual("")

	// check ctrl.Show
	req, err = http.NewRequest("GET", domain+path+"/"+id, strings.NewReader(""))
	mockstar.Expect(err).ToBeNil()
	req.Header.Set("API-Key", key)

	router.ServeHTTP(writer, req)

	call = ctrl.GetCall("Show", mockstar.Any, mockstar.Any)
	count = call.GetCallCount()
	mockstar.Expect(count).ToEqual(1)
	args = call.GetCallParams(0)
	resp, ok = args.Get(0).(Response)
	mockstar.Expect(ok).ToBeTrue()

	mockstar.Expect(writer.HasCalled("WriteHeader", http.StatusOK).Twice()).ToBeTrue()
	resp.Respond()
	mockstar.Expect(writer.HasCalled("WriteHeader", http.StatusOK).Times(3)).ToBeTrue()

	reqObj, ok = args.Get(1).(Request)
	mockstar.Expect(ok).ToBeTrue()
	mockstar.Expect(reqObj.Request()).ToEqual(req)
	mockstar.Expect(reqObj.Var("id")).ToEqual(id)

	// check ctrl.Destroy
	req, err = http.NewRequest("DELETE", domain+path+"/"+id, strings.NewReader(""))
	mockstar.Expect(err).ToBeNil()
	req.Header.Set("API-Key", key)

	router.ServeHTTP(writer, req)

	call = ctrl.GetCall("Destroy", mockstar.Any, mockstar.Any)
	count = call.GetCallCount()
	mockstar.Expect(count).ToEqual(1)
	args = call.GetCallParams(0)
	resp, ok = args.Get(0).(Response)
	mockstar.Expect(ok).ToBeTrue()

	mockstar.Expect(writer.HasCalled("WriteHeader", http.StatusOK).Times(3)).ToBeTrue()
	resp.Respond()
	mockstar.Expect(writer.HasCalled("WriteHeader", http.StatusOK).Times(4)).ToBeTrue()

	reqObj, ok = args.Get(1).(Request)
	mockstar.Expect(ok).ToBeTrue()
	mockstar.Expect(reqObj.Request()).ToEqual(req)
	mockstar.Expect(reqObj.Var("id")).ToEqual(id)
}

func Test_Router_Resource_Patchable(t *testing.T) {
	mockstar.T = t

	ctrl := newMockPatchable()
	ctrl.When("Apply", mockstar.Any, mockstar.Any).Return()

	writer := gophermocks.NewMockResponseWriter()
	writer.When("WriteHeader", http.StatusOK).Return()

	key := "This_is_a_key"
	path := "/mock"
	id := "This_is_the_id"

	router := NewRouter(map[Key]KeyConfig{
		Key(key): KeyConfig{true, []string{}},
	}, false)

	router.Resource(path, ctrl)

	domain := "http://1.2.3.4:5678"

	// check ctrl.Apply
	req, err := http.NewRequest("PATCH", domain+path+"/"+id, strings.NewReader(""))
	mockstar.Expect(err).ToBeNil()
	req.Header.Set("API-Key", key)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(writer, req)

	call := ctrl.GetCall("Apply", mockstar.Any, mockstar.Any)
	count := call.GetCallCount()
	mockstar.Expect(count).ToEqual(1)
	args := call.GetCallParams(0)
	resp, ok := args.Get(0).(Response)
	mockstar.Expect(ok).ToBeTrue()

	mockstar.Expect(writer.HasCalled("WriteHeader", http.StatusOK).Times(0)).ToBeTrue()
	resp.Respond()
	mockstar.Expect(writer.HasCalled("WriteHeader", http.StatusOK).Once()).ToBeTrue()

	reqObj, ok := args.Get(1).(Request)
	mockstar.Expect(ok).ToBeTrue()
	mockstar.Expect(reqObj.Request()).ToEqual(req)
	mockstar.Expect(reqObj.Var("id")).ToEqual(id)

}

func Test_Router_Resource_Updateable(t *testing.T) {
	mockstar.T = t

	ctrl := newMockUpdateable()
	ctrl.When("Update", mockstar.Any, mockstar.Any).Return()

	writer := gophermocks.NewMockResponseWriter()
	writer.When("WriteHeader", http.StatusOK).Return()

	key := "This_is_a_key"
	path := "/mock"
	id := "This_is_the_id"

	router := NewRouter(map[Key]KeyConfig{
		Key(key): KeyConfig{true, []string{}},
	}, false)

	router.Resource(path, ctrl)

	domain := "http://1.2.3.4:5678"

	// check ctrl.Apply
	req, err := http.NewRequest("PUT", domain+path+"/"+id, strings.NewReader(""))
	mockstar.Expect(err).ToBeNil()
	req.Header.Set("API-Key", key)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(writer, req)

	call := ctrl.GetCall("Update", mockstar.Any, mockstar.Any)
	count := call.GetCallCount()
	mockstar.Expect(count).ToEqual(1)
	args := call.GetCallParams(0)
	resp, ok := args.Get(0).(Response)
	mockstar.Expect(ok).ToBeTrue()

	mockstar.Expect(writer.HasCalled("WriteHeader", http.StatusOK).Times(0)).ToBeTrue()
	resp.Respond()
	mockstar.Expect(writer.HasCalled("WriteHeader", http.StatusOK).Once()).ToBeTrue()

	reqObj, ok := args.Get(1).(Request)
	mockstar.Expect(ok).ToBeTrue()
	mockstar.Expect(reqObj.Request()).ToEqual(req)
	mockstar.Expect(reqObj.Var("id")).ToEqual(id)

}

type mockController struct {
	*mockstar.Mock
}

func newMockController() *mockController {
	return &mockController{mockstar.NewMock()}
}

func (m *mockController) Index(resp Response, req Request) {
	m.Mock.Called("Index", resp, req)
}

func (m *mockController) Store(resp Response, req Request) {
	m.Mock.Called("Store", resp, req)
}

func (m *mockController) Show(resp Response, req Request) {
	m.Mock.Called("Show", resp, req)
}

func (m *mockController) Destroy(resp Response, req Request) {
	m.Mock.Called("Destroy", resp, req)
}

type mockUpdateable struct {
	*mockController
}

func newMockUpdateable() *mockUpdateable {
	return &mockUpdateable{newMockController()}
}

func (m *mockUpdateable) Update(resp Response, req Request) {
	m.Mock.Called("Update", resp, req)
}

type mockPatchable struct {
	*mockController
}

func newMockPatchable() *mockPatchable {
	return &mockPatchable{newMockController()}
}

func (m *mockPatchable) Apply(resp Response, req Request) {
	m.Mock.Called("Apply", resp, req)
}
