package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestIsPost_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		Context: &Context{
			Request: &http.Request{
				Method: "POST",
			},
		},
	}
	if !input.IsPost() {
		t.Errorf("IsPost() = %v, want %v", input.IsPost(), true)
	}
}
