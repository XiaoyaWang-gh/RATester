package fiber

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestShutdownWithContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	app := New()

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()

	err := app.ShutdownWithContext(ctx)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
