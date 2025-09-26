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
		},
		&route{
			matchers: matchersTree{
				matcher: func(ConnData) bool { return false },
			},
		},
	}
	r.Swap(0, 1)
	if r[0].matchers.matcher(ConnData{}) != false {
		t.Errorf("Swap failed")
	}
	if r[1].matchers.matcher(ConnData{}) != true {
		t.Errorf("Swap failed")
	}
}
