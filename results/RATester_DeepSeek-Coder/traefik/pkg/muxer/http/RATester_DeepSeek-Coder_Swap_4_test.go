package http

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSwap_4(t *testing.T) {
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
			handler:  nil,
			priority: 1,
		},
		&route{
			matchers: matchersTree{
				matcher: func(req *http.Request) bool {
					return req.Method == "POST"
				},
			},
			handler:  nil,
			priority: 2,
		},
	}

	r.Swap(0, 1)

	if r[0].priority != 2 || r[1].priority != 1 {
		t.Errorf("Expected priorities to be swapped, got %d and %d", r[0].priority, r[1].priority)
	}
}
