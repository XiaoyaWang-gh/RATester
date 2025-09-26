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
		t.Errorf("IsGet() = %v, want %v", input.IsGet(), true)
	}
}
