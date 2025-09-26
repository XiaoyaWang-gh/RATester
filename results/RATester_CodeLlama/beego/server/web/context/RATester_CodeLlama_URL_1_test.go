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
					Path: "/path/to/resource",
				},
			},
		},
	}
	if input.URL() != "/path/to/resource" {
		t.Errorf("URL() = %v, want %v", input.URL(), "/path/to/resource")
	}
}
