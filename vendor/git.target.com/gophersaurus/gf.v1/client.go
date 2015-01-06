package gf

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

var Client HttpClient = &httpClient{}

type HttpClient interface {
	NewRequest(method, url string, body []byte) (ClientRequest, error)
	Send(req ClientRequest) ([]byte, error)
}

type ClientRequest interface {
	Request() *http.Request
	AddHeader(name, value string)
	SetHeader(name, value string)
}

type httpClient struct {
}

func (c *httpClient) NewRequest(method, url string, body []byte) (ClientRequest, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	return &clientRequest{req}, nil
}

func (c *httpClient) Send(req ClientRequest) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req.Request())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	return bytes, err
}

type clientRequest struct {
	req *http.Request
}

func (c *clientRequest) Request() *http.Request {
	return c.req
}

func (c *clientRequest) AddHeader(name, value string) {
	c.req.Header.Add(name, value)

}

func (c *clientRequest) SetHeader(name, value string) {
	c.req.Header.Set(name, value)
}
