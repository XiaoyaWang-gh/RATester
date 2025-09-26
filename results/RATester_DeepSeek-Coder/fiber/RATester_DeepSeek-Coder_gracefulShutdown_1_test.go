package fiber

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGracefulShutdown_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	app := New()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := ListenConfig{
		GracefulContext: ctx,
		OnShutdownError: func(err error) {
			if err != nil {
				t.Errorf("Expected nil error, got %v", err)
			}
		},
		OnShutdownSuccess: func() {
			t.Log("Shutdown successful")
		},
	}

	go func() {
		err := app.Listen(":3000", cfg)
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}
	}()

	time.Sleep(100 * time.Millisecond)

	cancel()

	time.Sleep(100 * time.Millisecond)
}
