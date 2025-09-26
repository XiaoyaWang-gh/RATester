package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestDomain_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		Context: &Context{
			Request: &http.Request{
				Host: "example.com",
			},
		},
	}

	expected := "example.com"
	actual := input.Domain()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
