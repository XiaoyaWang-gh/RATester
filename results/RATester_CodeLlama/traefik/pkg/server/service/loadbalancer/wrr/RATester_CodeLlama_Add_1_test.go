package wrr

import (
	"fmt"
	"testing"
)

func TestAdd_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &Balancer{}
	b.Add("name", nil, nil)
	if len(b.handlers) != 1 {
		t.Errorf("expected 1 handler, got %d", len(b.handlers))
	}
}
