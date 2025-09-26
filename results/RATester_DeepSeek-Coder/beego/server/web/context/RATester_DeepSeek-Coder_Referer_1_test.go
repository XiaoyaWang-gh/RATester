package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestReferer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		Context: &Context{
			Request: &http.Request{
				Header: http.Header{
					"Referer": []string{"http://example.com"},
				},
			},
		},
	}

	referer := input.Referer()
	if referer != "http://example.com" {
		t.Errorf("Expected 'http://example.com', got '%s'", referer)
	}
}
