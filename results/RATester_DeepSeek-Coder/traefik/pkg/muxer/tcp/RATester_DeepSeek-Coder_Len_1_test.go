package tcp

import (
	"fmt"
	"testing"
)

func TestLen_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := routes{
		&route{
			matchers: matchersTree{
				matcher: func(ConnData) bool { return true },
			},
			handler: nil,
		},
		&route{
			matchers: matchersTree{
				matcher: func(ConnData) bool { return true },
			},
			handler: nil,
		},
	}

	if r.Len() != 2 {
		t.Errorf("Expected length of routes to be 2, got %d", r.Len())
	}
}
