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
					return req.Method == "GET"
				},
			},
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, "Hello, world")
			}),
		},
		&route{
			matchers: matchersTree{
				matcher: func(req *http.Request) bool {
					return req.Method == "POST"
				},
			},
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, "Hello, world")
			}),
		},
	}

	if r.Len() != 2 {
		t.Errorf("Expected length of routes to be 2, got %d", r.Len())
	}
}
