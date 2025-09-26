package mock

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestMock_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mockResponseFilter := &MockResponseFilter{}

	resp := &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Proto:      "HTTP/1.0",
		ProtoMajor: 1,
		ProtoMinor: 0,
		Header:     make(http.Header),
		Body:       io.NopCloser(strings.NewReader("")),
	}

	mockResponseFilter.Mock(nil, resp, nil)

	if len(mockResponseFilter.ms) != 1 {
		t.Errorf("Expected 1 mock, got %d", len(mockResponseFilter.ms))
	}
}
