package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestIsPut_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		Context: &Context{
			Request: &http.Request{
				Method: "PUT",
			},
		},
	}
	if !input.IsPut() {
		t.Errorf("IsPut() = %v, want %v", input.IsPut(), true)
	}
}
