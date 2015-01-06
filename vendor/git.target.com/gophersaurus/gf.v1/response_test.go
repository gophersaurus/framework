package gf

import (
	"errors"
	"net/http"
	"testing"

	"git.target.com/gophersaurus/gophersaurus/vendor/git.target.com/gophersaurus/gf.v1/gophermocks"
	"git.target.com/gophersaurus/gophersaurus/vendor/git.target.com/gophersaurus/gf.v1/mockstar"
	"git.target.com/gophersaurus/gophersaurus/vendor/gopkg.in/unrolled/render.v1"
)

func Test_Respond(t *testing.T) {
	mockstar.T = t

	mockRW := gophermocks.NewMockResponseWriter()
	mockRW.When("WriteHeader", http.StatusOK).Return()

	resp := buildResponse(mockRW)
	resp.Respond()

	mockstar.Expect(mockRW.HasCalled("WriteHeader", http.StatusOK).Once()).ToBeTrue()
}

func Test_RespondWithCode(t *testing.T) {
	mockstar.T = t

	mockRW := gophermocks.NewMockResponseWriter()
	mockRW.When("WriteHeader", http.StatusNoContent).Return()

	resp := buildResponse(mockRW)
	resp.HttpStatus(http.StatusNoContent)
	resp.Respond()

	mockstar.Expect(mockRW.HasCalled("WriteHeader", http.StatusNoContent).Once()).ToBeTrue()
}

func Test_RespondWithJson(t *testing.T) {
	mockstar.T = t

	renderer = render.New(render.Options{
		IndentJSON: false,
	})

	header := http.Header{}

	mockRW := gophermocks.NewMockResponseWriter()
	mockRW.When("WriteHeader", http.StatusCreated).Return()
	expectedBody := []byte("{\"a\":\"1\",\"b\":\"2\",\"c\":\"3\"}")
	mockRW.When("Write", expectedBody).Return(1, nil)
	mockRW.When("Header").Return(header)

	body := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}

	resp := buildResponse(mockRW)
	resp.HttpStatus(http.StatusCreated)
	resp.RespondWithJSON(body)

	mockstar.Expect(mockRW.HasCalled("WriteHeader", http.StatusCreated).Once()).ToBeTrue()
	mockstar.Expect(mockRW.HasCalled("Write", expectedBody).Once()).ToBeTrue()

	mockstar.Expect(len(header)).ToEqual(1)
	mockstar.Expect(header.Get("Content-Type")).ToEqual("application/json; charset=UTF-8")
}

func Test_RespondWithErr(t *testing.T) {
	mockstar.T = t

	renderer = render.New(render.Options{
		IndentJSON: false,
	})

	header := http.Header{}

	mockRW := gophermocks.NewMockResponseWriter()
	mockRW.When("WriteHeader", http.StatusInternalServerError).Return()
	expectedErr := []byte("{\"error\":\"%v\"}")
	mockRW.When("Write", mockstar.Any).Return(1, nil)
	mockRW.When("Header").Return(header)

	message := "This is an error"

	resp := buildResponse(mockRW)
	resp.RespondWithErr(errors.New(message))

	mockstar.Expect(mockRW.HasCalled("WriteHeader", http.StatusInternalServerError).Once()).ToBeTrue()
	mockstar.Expect(mockRW.HasCalled("Write", expectedErr).Once()).ToBeTrue()

	mockstar.Expect(len(header)).ToEqual(1)
	mockstar.Expect(header.Get("Content-Type")).ToEqual("application/json; charset=UTF-8")
}
