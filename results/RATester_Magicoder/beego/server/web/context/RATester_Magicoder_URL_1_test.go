package context

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestURL_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		Context: &Context{
			Request: &http.Request{
				URL: &url.URL{
					Path: "/test",
				},
			},
		},
	}

	expected := "/test"
	result := input.URL()

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
