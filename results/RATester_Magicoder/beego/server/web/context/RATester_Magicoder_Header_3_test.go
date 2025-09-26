package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHeader_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		Context: &Context{
			Request: &http.Request{
				Header: http.Header{
					"Content-Type": []string{"application/json"},
				},
			},
		},
	}

	header := input.Header("Content-Type")
	if header != "application/json" {
		t.Errorf("Expected 'application/json', got '%s'", header)
	}
}
