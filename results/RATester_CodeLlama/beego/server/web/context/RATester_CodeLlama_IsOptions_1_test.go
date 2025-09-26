package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestIsOptions_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		Context: &Context{
			Request: &http.Request{
				Method: "OPTIONS",
			},
		},
	}
	if !input.IsOptions() {
		t.Errorf("IsOptions() = %v, want %v", input.IsOptions(), true)
	}
}
