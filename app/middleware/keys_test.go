package middleware

import (
	"net/http"
	"strings"
	"testing"

	"git.target.com/gophersaurus/gf.v1"
	"git.target.com/gophersaurus/gf.v1/mock/gophermocks"
	"git.target.com/gophersaurus/gf.v1/mock/mockstar"
)

func Test_KeyHandler_NoKeyInRequest(t *testing.T) {
	mockstar.T = t

	domain := "1.2.3.4:5678"
	url := domain + "/path"

	key := "This_is_a_key"
	keyObj := gf.Key(key)
	keyConf := gf.KeyConfig{true, []string{}}
	keyMap := map[gf.Key]gf.KeyConfig{keyObj: keyConf}
	handler := NewKeyHandler(keyMap)

	writer := gophermocks.NewMockResponseWriter()
	responseHeader := http.Header{}
	writer.When("Header").Return(responseHeader)
	expectedErr := []byte("{\"error\":\"invalid permission\",\"status\":\"Forbidden\"}")
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
	keyObj := gf.Key(key)
	keyConf := gf.KeyConfig{true, []string{}}
	keyMap := map[gf.Key]gf.KeyConfig{keyObj: keyConf}
	handler := NewKeyHandler(keyMap)

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
	keyObj := gf.Key(key)
	keyConf := gf.KeyConfig{true, []string{}}
	keyMap := map[gf.Key]gf.KeyConfig{keyObj: keyConf}
	handler := NewKeyHandler(keyMap)

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
	keyObj := gf.Key(goodKey)
	keyConf := gf.KeyConfig{true, []string{}}
	keyMap := map[gf.Key]gf.KeyConfig{keyObj: keyConf}
	handler := NewKeyHandler(keyMap)

	writer := gophermocks.NewMockResponseWriter()
	responseHeader := http.Header{}
	writer.When("Header").Return(responseHeader)
	expectedErr := []byte("{\"error\":\"invalid permission\",\"status\":\"Forbidden\"}")
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

	keyObj := gf.Key(key)
	keyConf := gf.KeyConfig{true, []string{"5.4.3.2:1111"}}
	keyMap := map[gf.Key]gf.KeyConfig{keyObj: keyConf}
	handler := NewKeyHandler(keyMap)

	writer := gophermocks.NewMockResponseWriter()
	responseHeader := http.Header{}
	writer.When("Header").Return(responseHeader)
	expectedErr := []byte("{\"error\":\"invalid permission\",\"status\":\"Forbidden\"}")
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

	keyObj := gf.Key(key)
	keyConf := gf.KeyConfig{true, []string{sender}}
	keyMap := map[gf.Key]gf.KeyConfig{keyObj: keyConf}
	handler := NewKeyHandler(keyMap)

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

func Test_Router_Get_BadKey(t *testing.T) {
	mockstar.T = t

	ctrl := gophermocks.NewMockController()
	ctrl.When("Index", mockstar.Any, mockstar.Any).Return()

	writer := gophermocks.NewMockResponseWriter()
	responseHeader := http.Header{}
	writer.When("Header").Return(responseHeader)
	writer.When("WriteHeader", http.StatusForbidden).Return()
	expectedErr := []byte("{\"error\":\"invalid permission\",\"status\":\"Forbidden\"}")
	writer.When("Write", expectedErr).Return(len(expectedErr), nil)

	key := "This_is_a_key"
	path := "/mock"

	keyHandler := NewKeyHandler(map[gf.Key]gf.KeyConfig{
		gf.Key(key): gf.KeyConfig{true, []string{}},
	})

	router := gf.NewRouter()

	router.Get(path, ctrl.Index, keyHandler)

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
