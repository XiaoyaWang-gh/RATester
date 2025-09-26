package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCookie_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		Context: &Context{
			Request: &http.Request{
				Header: http.Header{
					"Cookie": []string{"key=value"},
				},
			},
		},
	}
	if input.Cookie("key") != "value" {
		t.Errorf("Cookie() = %v, want %v", input.Cookie("key"), "value")
	}
}
