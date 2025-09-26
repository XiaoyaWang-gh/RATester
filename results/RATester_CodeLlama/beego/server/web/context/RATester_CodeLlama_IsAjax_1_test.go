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
				Header: map[string][]string{
					"X-Requested-With": []string{"XMLHttpRequest"},
				},
			},
		},
	}
	if !input.IsAjax() {
		t.Error("IsAjax() should return true")
	}
}
