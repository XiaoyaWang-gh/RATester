package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestReferer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		Context: &Context{
			Request: &http.Request{
				Header: map[string][]string{
					"Referer": []string{"http://www.google.com"},
				},
			},
		},
	}
	if input.Referer() != "http://www.google.com" {
		t.Errorf("Referer() = %v, want %v", input.Referer(), "http://www.google.com")
	}
}
