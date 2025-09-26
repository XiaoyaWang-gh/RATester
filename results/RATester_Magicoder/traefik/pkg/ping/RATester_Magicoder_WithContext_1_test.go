package ping

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestWithContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	handler := &Handler{}
	handler.WithContext(ctx)

	// Simulate the context being canceled
	cancel()

	// Wait for the goroutine to finish
	time.Sleep(100 * time.Millisecond)

	if !handler.terminating {
		t.Error("Handler should have been terminating after context cancellation")
	}
}
