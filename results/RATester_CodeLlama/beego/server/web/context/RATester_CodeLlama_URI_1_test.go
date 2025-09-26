package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestURI_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		Context: &Context{
			Request: &http.Request{
				RequestURI: "/test",
			},
		},
	}
	if input.URI() != "/test" {
		t.Errorf("URI() = %v, want %v", input.URI(), "/test")
	}
}
