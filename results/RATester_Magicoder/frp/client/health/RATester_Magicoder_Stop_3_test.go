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
		ctx:    ctx,
		cancel: cancel,
	}

	monitor.Stop()

	if monitor.ctx.Err() == nil {
		t.Error("Expected context to be canceled after calling Stop(), but it is not.")
	}
}
