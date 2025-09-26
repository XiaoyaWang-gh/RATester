package try

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestResponseUntilStatusCode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	req := &http.Request{
		Method: "GET",
		URL: &url.URL{
			Scheme: "http",
			Host:   "example.com",
			Path:   "/path",
		},
	}
	timeout := time.Second * 10
	statusCode := 200

	// when
	response, err := ResponseUntilStatusCode(req, timeout, statusCode)

	// then
	if err != nil {
		t.Errorf("ResponseUntilStatusCode() error = %v", err)
		return
	}
	if response == nil {
		t.Errorf("ResponseUntilStatusCode() response = nil")
		return
	}
	if response.StatusCode != statusCode {
		t.Errorf("ResponseUntilStatusCode() response.StatusCode = %v, want %v", response.StatusCode, statusCode)
	}
}
