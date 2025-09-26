package fiber

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestgracefulShutdown_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{}
	ctx, cancel := context.WithCancel(context.Background())
	cfg := ListenConfig{
		GracefulContext: ctx,
		OnShutdownError: func(err error) {
			t.Errorf("Unexpected error: %v", err)
		},
		OnShutdownSuccess: func() {
			t.Errorf("Unexpected success")
		},
	}

	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()

	app.gracefulShutdown(ctx, cfg)
}
