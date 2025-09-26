package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestPort_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		Context: &Context{
			Request: &http.Request{
				Host: "localhost:8080",
			},
		},
	}

	port := input.Port()
	if port != 8080 {
		t.Errorf("Expected port 8080, got %d", port)
	}
}
