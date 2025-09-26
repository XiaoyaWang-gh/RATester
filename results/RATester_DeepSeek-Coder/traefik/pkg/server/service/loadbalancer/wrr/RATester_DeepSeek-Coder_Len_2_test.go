package wrr

import (
	"fmt"
	"testing"
)

func TestLen_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &Balancer{
		handlers: []*namedHandler{
			{name: "handler1"},
			{name: "handler2"},
			{name: "handler3"},
		},
	}

	length := b.Len()
	if length != 3 {
		t.Errorf("Expected length of 3, got %d", length)
	}
}
