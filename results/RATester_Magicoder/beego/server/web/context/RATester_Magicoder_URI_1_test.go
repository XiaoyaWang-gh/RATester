package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestURI_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		Context: &Context{
			Request: &http.Request{
				RequestURI: "/test",
			},
		},
	}

	expected := "/test"
	result := input.URI()

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
