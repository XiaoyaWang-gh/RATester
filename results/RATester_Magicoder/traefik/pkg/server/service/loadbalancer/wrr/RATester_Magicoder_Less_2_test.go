package wrr

import (
	"fmt"
	"testing"
)

func TestLess_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &Balancer{
		handlers: []*namedHandler{
			{name: "handler1", deadline: 1.0},
			{name: "handler2", deadline: 2.0},
			{name: "handler3", deadline: 3.0},
		},
	}

	if !b.Less(0, 1) {
		t.Error("Expected handler1 to be less than handler2")
	}

	if b.Less(1, 0) {
		t.Error("Expected handler2 not to be less than handler1")
	}

	if !b.Less(1, 2) {
		t.Error("Expected handler2 to be less than handler3")
	}

	if b.Less(2, 1) {
		t.Error("Expected handler3 not to be less than handler2")
	}
}
