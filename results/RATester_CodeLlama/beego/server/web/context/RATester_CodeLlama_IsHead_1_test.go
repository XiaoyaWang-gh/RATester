package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestIsHead_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		Context: &Context{
			Request: &http.Request{
				Method: "HEAD",
			},
		},
	}
	if !input.IsHead() {
		t.Errorf("IsHead() = %v, want %v", input.IsHead(), true)
	}
}
