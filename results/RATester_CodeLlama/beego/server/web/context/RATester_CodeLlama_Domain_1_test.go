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
				Host: "localhost",
			},
		},
	}
	if input.Domain() != "localhost" {
		t.Errorf("Domain() = %v, want %v", input.Domain(), "localhost")
	}
}
