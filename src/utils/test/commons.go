package test

import (
	"bytes"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type HandlerTestCase struct {
	Name         string
	ExpectedCode int
	ExpectedBody string
}

type HandlerTest struct {
	Router   *mux.Router
	Writer   *httptest.ResponseRecorder
	Req      *http.Request
	TestCase HandlerTestCase
	T        *testing.T
}

func CheckBody(body *bytes.Buffer, expectedBody *string) bool {
	return (body != nil && expectedBody != nil && strings.TrimSuffix(body.String(), "\n") == *expectedBody) ||
		(body == nil && expectedBody == nil)
}

func DoHandlerTest(test *HandlerTest) {
	test.Router.ServeHTTP(test.Writer, test.Req)

	if status := test.Writer.Code; status != test.TestCase.ExpectedCode {
		test.T.Errorf("Incorrect status code: got %v want %v", status, test.TestCase.ExpectedCode)
	} else if !CheckBody(test.Writer.Body, &test.TestCase.ExpectedBody) {
		test.T.Errorf("Incorrect body: got %v want %v", test.Writer.Body.String(), test.TestCase.ExpectedBody)
	}
}
