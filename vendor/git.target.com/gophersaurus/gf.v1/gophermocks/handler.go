package gophermocks

import (
	"net/http"

	"git.target.com/gophersaurus/gophersaurus/vendor/git.target.com/gophersaurus/gf.v1/mockstar"
)

type mockHandler struct {
	*mockstar.Mock
}

func NewMockHandler() *mockHandler {
	return &mockHandler{mockstar.NewMock()}
}

func (m *mockHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	m.Mock.Called("ServeHTTP", rw, r)
}
