package context

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestScheme_1(t *testing.T) {
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
				},
			},
		},
	}
	if input.Scheme() != "http" {
		t.Errorf("Scheme() = %v, want %v", input.Scheme(), "http")
	}
}
