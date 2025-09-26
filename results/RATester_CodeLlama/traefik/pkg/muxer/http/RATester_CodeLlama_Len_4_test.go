package http

import (
	"fmt"
	"net/http"
	"testing"
)

func TestLen_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := routes{
		&route{
			matchers: matchersTree{
				matcher: func(req *http.Request) bool {
					return true
				},
			},
		},
	}
	if r.Len() != 1 {
		t.Errorf("Expected 1, got %d", r.Len())
	}
}
