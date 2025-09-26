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

	b := &Balancer{}
	b.handlers = []*namedHandler{
		{name: "a"},
		{name: "b"},
		{name: "c"},
	}
	if got := b.Len(); got != 3 {
		t.Errorf("Len() = %v, want %v", got, 3)
	}
}
