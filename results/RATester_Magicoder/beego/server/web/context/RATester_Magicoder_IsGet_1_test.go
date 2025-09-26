package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestIsGet_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		Context: &Context{
			Request: &http.Request{
				Method: "GET",
			},
		},
	}

	if !input.IsGet() {
		t.Error("Expected IsGet to return true for GET request")
	}

	input.Context.Request.Method = "POST"
	if input.IsGet() {
		t.Error("Expected IsGet to return false for non-GET request")
	}
}
