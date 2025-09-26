package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHost_1(t *testing.T) {
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
	if input.Host() != "localhost" {
		t.Errorf("Host() = %v, want %v", input.Host(), "localhost")
	}
}
