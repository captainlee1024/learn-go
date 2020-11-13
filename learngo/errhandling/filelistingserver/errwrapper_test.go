package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errPanic(writer http.ResponseWriter,
	request *http.Request) error {
	panic(123)
}

func errUserError(writer http.ResponseWriter,
	request *http.Request) error {
	return testinguserError("user error")
}

func errNotFound(writer http.ResponseWriter,
	request *http.Request) error {
	return os.ErrNotExist
}

func errNoPermission(writer http.ResponseWriter,
	request *http.Request) error {
	return os.ErrPermission
}

func errUnknow(writer http.ResponseWriter,
	request *http.Request) error {
	return errors.New("unknow error")
}

func noError(writer http.ResponseWriter,
	request *http.Request) error {
	fmt.Fprintln(writer, "no error")
	return nil
}

type testinguserError string

func (e testinguserError) Error() string {
	return e.Message()
}

func (e testinguserError) Message() string {
	return string(e)
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	// 在这里写函数太长，我们把函数提出去
	{h: errPanic, code: 500, message: "Internal Server Error"},
	{h: errUserError, code: 400, message: "user error"},
	{errNotFound, 404, "Not Found"},
	{errNoPermission, 403, "Forbidden"},
	{errUnknow, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

func verifyResponse(response *http.Response,
	expectedCode int, expectedMessage string, t *testing.T) {
	// 返回的 error 我们不需要使用
	b, _ := ioutil.ReadAll(response.Body)
	// body := string(b)
	//把最后的换行去掉
	body := strings.Trim(string(b), "\n")

	if response.StatusCode != expectedCode ||
		body != expectedMessage {
		t.Errorf("expect (%d %s);"+
			"got (%d %s)", expectedCode, expectedMessage, response.StatusCode, body)
	}
}

func TestErrWrapper(t *testing.T) {
	// tests := []struct {
	// 	h       appHandler
	// 	code    int
	// 	message string
	// }{
	// 	// 在这里写函数太长，我们把函数提出去
	// 	{h: errPanic, code: 500, message: "Internal Server Error"},
	// 	{h: errUserError, code: 400, message: "user error"},
	// 	{errNotFound, 404, "Not Found"},
	// 	{errNoPermission, 403, "Forbidden"},
	// 	{errUnknow, 500, "Internal Server Error"},
	// 	{noError, 200, "no error"},
	// }

	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"http://www.baidu.com",
			nil,
		)

		f(response, request)

		// // 返回的 error 我们不需要使用
		// b, _ := ioutil.ReadAll(response.Body)
		// // body := string(b)
		// //把最后的换行去掉
		// body := strings.Trim(string(b), "\n")

		// if response.Code != tt.code ||
		// 	body != tt.message {
		// 	t.Errorf("expect (%d %s);"+
		// 		"got (%d %s)", tt.code, tt.message, response.Code, body)
		// }

		verifyResponse(response.Result(), tt.code, tt.message, t)
	}
}

func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)
		verifyResponse(resp, tt.code, tt.message, t)
	}
}
