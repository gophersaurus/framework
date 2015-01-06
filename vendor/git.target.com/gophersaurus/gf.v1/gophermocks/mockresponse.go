package gophermocks

import (
	"fmt"
	"net/http"

	"git.target.com/gophersaurus/gophersaurus/vendor/git.target.com/gophersaurus/gf.v1/mockstar"
)

type mockResponseWriter struct {
	*mockstar.Mock
}

func NewMockResponseWriter() *mockResponseWriter {
	return &mockResponseWriter{mockstar.NewMock()}
}

func (m *mockResponseWriter) Header() http.Header {
	args := m.Mock.Called("Header")
	return m.castHeader(args.Get(0))
}

func (m *mockResponseWriter) castHeader(obj interface{}) http.Header {
	if obj == nil {
		return nil
	}
	out, ok := obj.(http.Header)
	if !ok {
		mockstar.Fatal("cannot cast value to http.Header: " + fmt.Sprintf("%v", obj))
	}
	return out
}

func (m *mockResponseWriter) Write(data []byte) (int, error) {
	args := m.Mock.Called("Write", data)
	return args.Int(0), args.Err(1)
}

func (m *mockResponseWriter) WriteHeader(statusCode int) {
	m.Mock.Called("WriteHeader", statusCode)
}
