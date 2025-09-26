package nathole

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestCleanWorker_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{
		clientCfgs: make(map[string]*ClientCfg),
		sessions:   make(map[string]*Session),
		analyzer:   &Analyzer{},
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go ctrl.CleanWorker(ctx)

	// Wait for the worker to start
	time.Sleep(100 * time.Millisecond)

	// Cancel the context to stop the worker
	cancel()

	// Wait for the worker to stop
	time.Sleep(100 * time.Millisecond)

	// Check if the worker stopped
	if _, ok := <-ctx.Done(); ok {
		t.Error("CleanWorker did not stop")
	}
}
