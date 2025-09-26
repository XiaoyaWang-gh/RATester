package httplib

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSetCookie_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := &BeegoHTTPRequest{
		req: &http.Request{
			Header: make(http.Header),
		},
	}

	cookie := &http.Cookie{
		Name:  "test",
		Value: "value",
	}

	req.SetCookie(cookie)

	expected := "test=value"
	actual := req.req.Header.Get("Cookie")

	if actual != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}
