package gophermocks

import "git.target.com/gophersaurus/gophersaurus/vendor/git.target.com/gophersaurus/gf.v1/mockstar"

type mockResponse struct {
	*mockstar.Mock
}

func NewMockResponse() *mockResponse {
	return &mockResponse{mockstar.NewMock()}
}

func (m *mockResponse) RespondWithJSON(json map[string]string) {
	m.Mock.Called("RespondWithJSON", json)
}

func (m *mockResponse) Respond() {
	m.Mock.Called("Respond")
}

func (m *mockResponse) RespondWithErr(err error) {
	m.Mock.Called("RespondWithErr", err)
}

func (m *mockResponse) HttpStatus(code int) {
	m.Mock.Called("HttpStatus", code)
}

func (m *mockResponse) Body(obj interface{}) {
	m.Mock.Called("Body", obj)
}

func (m *mockResponse) AppendBody(objects ...interface{}) {
	m.Mock.CalledVarArgs("AppendBody", objects)
}

func (m *mockResponse) Header(key, value string) {
	m.Mock.Called("Header", key, value)
}

func (m *mockResponse) FlushHeaders() {
	m.Mock.Called("FlushHeaders")
}
