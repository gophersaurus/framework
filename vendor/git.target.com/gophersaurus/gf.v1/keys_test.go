package gf

import (
	"net/http"
	"strings"
	"testing"

	"git.target.com/gophersaurus/gophersaurus/vendor/git.target.com/gophersaurus/gf.v1/gophermocks"
	"git.target.com/gophersaurus/gophersaurus/vendor/git.target.com/gophersaurus/gf.v1/mockstar"
	"git.target.com/gophersaurus/gophersaurus/vendor/gopkg.in/unrolled/render.v1"
)

func Test_KeyHandler_NoKeyInRequest(t *testing.T) {
	mockstar.T = t

	domain := "1.2.3.4:5678"
	url := domain + "/path"

	key := "This_is_a_key"
	keyObj := Key(key)
	keyConf := KeyConfig{true, []string{}}
	keyMap := map[Key]KeyConfig{keyObj: keyConf}
	handler := &keyHandler{keyMap}

	renderer = render.New(render.Options{
		IndentJSON: false,
	})

	writer := gophermocks.NewMockResponseWriter()
	responseHeader := http.Header{}
	writer.When("Header").Return(responseHeader)
	expectedErr := []byte("{\"error\":\"invalid permissions\"}")
	writer.When("Write", expectedErr).Return(len(expectedErr), nil)
	writer.When("WriteHeader", http.StatusForbidden).Return()
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	mockstar.Expect(err).ToBeNil()
	next := gophermocks.NewMockHandler()
	next.When("ServeHTTP", writer, req).Return()

	handler.ServeHTTP(writer, req, next.ServeHTTP)

	mockstar.Expect(next.HasCalled("ServeHTTP", writer, req).Times(0)).ToBeTrue()

	mockstar.Expect(writer.HasCalled("Write", expectedErr).Once()).ToBeTrue()

	mockstar.Expect(writer.HasCalled("WriteHeader", http.StatusForbidden).Once()).ToBeTrue()

	mockstar.Expect(len(responseHeader)).ToEqual(1)
	mockstar.Expect(responseHeader.Get("Content-Type")).ToEqual("application/json; charset=UTF-8")
}

func Test_KeyHandler_SimpleKeyInQuery(t *testing.T) {
	mockstar.T = t

	domain := "1.2.3.4:5678"
	key := "This_is_a_key"
	url := domain + "/path?key=" + key
	keyObj := Key(key)
	keyConf := KeyConfig{true, []string{}}
	keyMap := map[Key]KeyConfig{keyObj: keyConf}
	handler := &keyHandler{keyMap}

	writer := gophermocks.NewMockResponseWriter()
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	mockstar.Expect(err).ToBeNil()
	next := gophermocks.NewMockHandler()
	next.When("ServeHTTP", writer, req).Return()

	handler.ServeHTTP(writer, req, next.ServeHTTP)

	mockstar.Expect(next.HasCalled("ServeHTTP", writer, req).Once()).ToBeTrue()
}

func Test_KeyHandler_SimpleKeyInHeader(t *testing.T) {
	mockstar.T = t

	domain := "1.2.3.4:5678"
	key := "This_is_a_key"
	url := domain + "/path"
	keyObj := Key(key)
	keyConf := KeyConfig{true, []string{}}
	keyMap := map[Key]KeyConfig{keyObj: keyConf}
	handler := &keyHandler{keyMap}

	writer := gophermocks.NewMockResponseWriter()
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	mockstar.Expect(err).ToBeNil()
	next := gophermocks.NewMockHandler()
	next.When("ServeHTTP", writer, req).Return()
	req.Header.Set("API-Key", key)

	handler.ServeHTTP(writer, req, next.ServeHTTP)

	mockstar.Expect(next.HasCalled("ServeHTTP", writer, req).Once()).ToBeTrue()
}

func Test_KeyHandler_SimpleKeyInHeaderNotInConfig(t *testing.T) {
	mockstar.T = t

	domain := "1.2.3.4:5678"
	url := domain + "/path"
	badKey := "This_key_is_NOT_in_the_config"

	goodKey := "This_key_is_in_the_config"
	keyObj := Key(goodKey)
	keyConf := KeyConfig{true, []string{}}
	keyMap := map[Key]KeyConfig{keyObj: keyConf}
	handler := &keyHandler{keyMap}

	renderer = render.New(render.Options{
		IndentJSON: false,
	})

	writer := gophermocks.NewMockResponseWriter()
	responseHeader := http.Header{}
	writer.When("Header").Return(responseHeader)
	expectedErr := []byte("{\"error\":\"invalid permissions\"}")
	writer.When("Write", expectedErr).Return(len(expectedErr), nil)
	writer.When("WriteHeader", http.StatusForbidden).Return()
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	mockstar.Expect(err).ToBeNil()
	req.Header.Set("API-Key", badKey)
	next := gophermocks.NewMockHandler()
	next.When("ServeHTTP", writer, req).Return()

	handler.ServeHTTP(writer, req, next.ServeHTTP)

	mockstar.Expect(next.HasCalled("ServeHTTP", writer, req).Times(0)).ToBeTrue()

	mockstar.Expect(writer.HasCalled("Write", expectedErr).Once()).ToBeTrue()

	mockstar.Expect(writer.HasCalled("WriteHeader", http.StatusForbidden).Once()).ToBeTrue()

	mockstar.Expect(len(responseHeader)).ToEqual(1)
	mockstar.Expect(responseHeader.Get("Content-Type")).ToEqual("application/json; charset=UTF-8")
}

func Test_KeyHandler_KeyInHeaderDomainNotInWhiteList(t *testing.T) {
	mockstar.T = t

	domain := "1.2.3.4:5678"
	url := domain + "/path"
	key := "This_is_a_key"

	sender := "5.6.7.8:1234"

	keyObj := Key(key)
	keyConf := KeyConfig{true, []string{"5.4.3.2:1111"}}
	keyMap := map[Key]KeyConfig{keyObj: keyConf}
	handler := &keyHandler{keyMap}

	renderer = render.New(render.Options{
		IndentJSON: false,
	})

	writer := gophermocks.NewMockResponseWriter()
	responseHeader := http.Header{}
	writer.When("Header").Return(responseHeader)
	expectedErr := []byte("{\"error\":\"invalid permissions\"}")
	writer.When("Write", expectedErr).Return(len(expectedErr), nil)
	writer.When("WriteHeader", http.StatusForbidden).Return()
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	mockstar.Expect(err).ToBeNil()
	req.Header.Set("API-Key", key)
	req.RemoteAddr = sender
	next := gophermocks.NewMockHandler()
	next.When("ServeHTTP", writer, req).Return()

	handler.ServeHTTP(writer, req, next.ServeHTTP)

	mockstar.Expect(next.HasCalled("ServeHTTP", writer, req).Times(0)).ToBeTrue()

	mockstar.Expect(writer.HasCalled("Write", expectedErr).Once()).ToBeTrue()

	mockstar.Expect(writer.HasCalled("WriteHeader", http.StatusForbidden).Once()).ToBeTrue()

	mockstar.Expect(len(responseHeader)).ToEqual(1)
	mockstar.Expect(responseHeader.Get("Content-Type")).ToEqual("application/json; charset=UTF-8")
}

func Test_KeyHandler_KeyInHeaderDomainInWhiteList(t *testing.T) {
	mockstar.T = t

	domain := "1.2.3.4:5678"
	url := domain + "/path"
	key := "This_is_a_key"

	sender := "5.6.7.8:1234"

	keyObj := Key(key)
	keyConf := KeyConfig{true, []string{sender}}
	keyMap := map[Key]KeyConfig{keyObj: keyConf}
	handler := &keyHandler{keyMap}

	renderer = render.New(render.Options{
		IndentJSON: false,
	})

	writer := gophermocks.NewMockResponseWriter()
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	mockstar.Expect(err).ToBeNil()
	req.Header.Set("API-Key", key)
	req.RemoteAddr = sender
	next := gophermocks.NewMockHandler()
	next.When("ServeHTTP", writer, req).Return()

	handler.ServeHTTP(writer, req, next.ServeHTTP)

	mockstar.Expect(next.HasCalled("ServeHTTP", writer, req).Once()).ToBeTrue()
}
