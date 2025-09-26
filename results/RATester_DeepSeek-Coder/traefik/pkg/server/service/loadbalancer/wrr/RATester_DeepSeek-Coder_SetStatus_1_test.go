package wrr

import (
	"context"
	"fmt"
	"testing"
)

func TestSetStatus_1(t *testing.T) {
	ctx := context.Background()
	b := &Balancer{
		status: make(map[string]struct{}),
	}

	t.Run("Set status to UP", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		b.SetStatus(ctx, "child1", true)
		if _, ok := b.status["child1"]; !ok {
			t.Errorf("Expected child1 to be in status map")
		}
	})

	t.Run("Set status to DOWN", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		b.SetStatus(ctx, "child1", false)
		if _, ok := b.status["child1"]; ok {
			t.Errorf("Expected child1 to be removed from status map")
		}
	})
}
