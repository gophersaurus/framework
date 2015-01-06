package mockclient

import (
	"fmt"
	"net/http"

	gf "git.target.com/gophersaurus/gophersaurus/vendor/git.target.com/gophersaurus/gf.v1"
	"git.target.com/gophersaurus/gophersaurus/vendor/git.target.com/gophersaurus/gf.v1/mockstar"
)

type mockClient struct {
	*mockstar.Mock
}

func NewMockClient() *mockClient {
	return &mockClient{mockstar.NewMock()}
}

func (c *mockClient) Send(req gf.ClientRequest) ([]byte, error) {
	args := c.Mock.Called("Send", req)
	return args.Bytes(0), args.Err(1)
}

func (c *mockClient) NewRequest(method, path string, body []byte) (gf.ClientRequest, error) {
	args := c.Mock.Called("NewRequest", method, path, body)
	return c.castClientRequest(args.Get(0)), args.Err(1)
}

func (c *mockClient) castClientRequest(obj interface{}) *mockClientRequest {
	if obj == nil {
		return nil
	}
	out, ok := obj.(*mockClientRequest)
	if !ok {
		mockstar.Fatal("cannot cast value to mockClientRequest: " + fmt.Sprintf("%v", obj))
	}
	return out
}

type mockClientRequest struct {
	*mockstar.Mock
}

func NewMockClientRequest() *mockClientRequest {
	return &mockClientRequest{mockstar.NewMock()}
}

func (c *mockClientRequest) Request() *http.Request {
	args := c.Mock.Called("Request")
	return c.castHttpRequest(args.Get(0))
}

func (c *mockClientRequest) castHttpRequest(obj interface{}) *http.Request {
	if obj == nil {
		return nil
	}
	out, ok := obj.(http.Request)
	if !ok {
		mockstar.Fatal("cannot cast value to http.Request: " + fmt.Sprintf("%v", obj))
	}
	return &out
}

func (c *mockClientRequest) AddHeader(name, value string) {
	c.Mock.Called("AddHeader", name, value)
}

func (c *mockClientRequest) SetHeader(name, value string) {
	c.Mock.Called("SetHeader", name, value)
}
