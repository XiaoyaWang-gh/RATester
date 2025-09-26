package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestAcceptsJSON_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		Context: &Context{
			Request: &http.Request{
				Header: http.Header{
					"Accept": []string{"application/json"},
				},
			},
		},
	}

	result := input.AcceptsJSON()
	if !result {
		t.Errorf("Expected true, got %v", result)
	}
}
