package consulcatalog

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestRepeatSend_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan struct{}, 1)

	go repeatSend(ctx, time.Second, ch)

	select {
	case <-ch:
		// Test passed
	case <-time.After(2 * time.Second):
		t.Error("Test timed out")
	}
}
