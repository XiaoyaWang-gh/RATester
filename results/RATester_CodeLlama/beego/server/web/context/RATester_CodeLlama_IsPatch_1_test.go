package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestIsPatch_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		Context: &Context{
			Request: &http.Request{
				Method: "PATCH",
			},
		},
	}
	if !input.IsPatch() {
		t.Errorf("IsPatch() = %v, want %v", input.IsPatch(), true)
	}
}
