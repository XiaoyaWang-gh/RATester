package fiber

import (
	"context"
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestShutdownWithContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{
		hooks: &Hooks{
			onShutdown: []func() error{
				func() error {
					return nil
				},
			},
		},
		server: &fasthttp.Server{},
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := app.ShutdownWithContext(ctx)
	if err != ErrNotRunning {
		t.Errorf("Expected error %v, got %v", ErrNotRunning, err)
	}

	app.server = &fasthttp.Server{}
	err = app.ShutdownWithContext(context.Background())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
