package context

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestIsSecure_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		Context: &Context{
			Request: &http.Request{
				URL: &url.URL{
					Scheme: "https",
				},
			},
		},
	}

	if !input.IsSecure() {
		t.Error("Expected IsSecure to return true for https scheme")
	}

	input.Context.Request.URL.Scheme = "http"
	if input.IsSecure() {
		t.Error("Expected IsSecure to return false for http scheme")
	}
}
