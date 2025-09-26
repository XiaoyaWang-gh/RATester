package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestIsAjax_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		Context: &Context{
			Request: &http.Request{
				Header: http.Header{
					"X-Requested-With": []string{"XMLHttpRequest"},
				},
			},
		},
	}

	if !input.IsAjax() {
		t.Error("Expected IsAjax to return true")
	}

	input.Context.Request.Header.Set("X-Requested-With", "NotXMLHttpRequest")

	if input.IsAjax() {
		t.Error("Expected IsAjax to return false")
	}
}
