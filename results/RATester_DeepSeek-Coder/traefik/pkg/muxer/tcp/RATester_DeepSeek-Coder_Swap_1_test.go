package tcp

import (
	"fmt"
	"testing"
)

func TestSwap_1(t *testing.T) {
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
			handler:  nil,
			catchAll: false,
			priority: 1,
		},
		&route{
			matchers: matchersTree{
				matcher: func(ConnData) bool { return false },
			},
			handler:  nil,
			catchAll: false,
			priority: 2,
		},
	}

	r.Swap(0, 1)

	if r[0].priority != 2 || r[1].priority != 1 {
		t.Errorf("Expected priorities to be swapped, got %v and %v", r[0].priority, r[1].priority)
	}
}
