package health

import (
	"context"
	"fmt"
	"testing"
)

func TestStop_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	monitor := &Monitor{
		checkType: "tcp",
		addr:      "localhost:8080",
		ctx:       ctx,
		cancel:    cancel,
	}

	monitor.Stop()

	if monitor.ctx.Err() != context.Canceled {
		t.Errorf("Expected context to be canceled, but it's not")
	}
}
