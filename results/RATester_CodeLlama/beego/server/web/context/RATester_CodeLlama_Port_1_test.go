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
	if input.Port() != 8080 {
		t.Errorf("Port() = %v, want %v", input.Port(), 8080)
	}
}
