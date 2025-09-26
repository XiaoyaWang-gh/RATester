package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestMethod_1(t *testing.T) {
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

	method := input.Method()

	if method != "GET" {
		t.Errorf("Expected method to be 'GET', but got '%s'", method)
	}
}
