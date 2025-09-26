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
			{
				name:     "a",
				weight:   1,
				deadline: 1,
			},
			{
				name:     "b",
				weight:   1,
				deadline: 2,
			},
		},
	}
	if !b.Less(0, 1) {
		t.Error("expected true")
	}
	if b.Less(1, 0) {
		t.Error("expected false")
	}
}
