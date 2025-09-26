package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHeader_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		Context: &Context{
			Request: &http.Request{
				Header: http.Header{
					"key": []string{"value"},
				},
			},
		},
	}
	if input.Header("key") != "value" {
		t.Errorf("Header() = %v, want %v", input.Header("key"), "value")
	}
}
