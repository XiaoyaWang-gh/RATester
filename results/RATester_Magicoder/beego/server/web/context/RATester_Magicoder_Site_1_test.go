package context

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestSite_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		Context: &Context{
			Request: &http.Request{
				URL: &url.URL{
					Scheme: "http",
					Host:   "example.com",
				},
			},
		},
	}

	expected := "http://example.com"
	result := input.Site()

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
