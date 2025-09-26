package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSubDomains_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		Context: &Context{
			Request: &http.Request{
				Host: "www.example.com",
			},
		},
	}

	result := input.SubDomains()
	expected := "www"

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
