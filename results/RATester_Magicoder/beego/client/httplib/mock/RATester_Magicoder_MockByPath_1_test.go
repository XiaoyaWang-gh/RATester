package mock

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestMockByPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mockResponseFilter := &MockResponseFilter{}
	path := "/test"
	resp := &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Proto:      "HTTP/1.0",
		ProtoMajor: 1,
		ProtoMinor: 0,
		Header:     make(http.Header),
		Body:       ioutil.NopCloser(strings.NewReader("test body")),
		Request:    &http.Request{Method: "GET", URL: &url.URL{Path: path}},
	}
	err := errors.New("test error")

	mockResponseFilter.MockByPath(path, resp, err)

	if len(mockResponseFilter.ms) != 1 {
		t.Errorf("Expected 1 mock, got %d", len(mockResponseFilter.ms))
	}

	mock := mockResponseFilter.ms[0]
	if mock.cond.(*SimpleCondition).path != path {
		t.Errorf("Expected path %s, got %s", path, mock.cond.(*SimpleCondition).path)
	}

	if mock.resp != resp {
		t.Errorf("Expected response %v, got %v", resp, mock.resp)
	}

	if mock.err != err {
		t.Errorf("Expected error %v, got %v", err, mock.err)
	}
}
