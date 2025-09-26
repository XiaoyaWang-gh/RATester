package safe

import (
	"context"
	"fmt"
	"testing"
)

func TestStop_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pool := &Pool{
		ctx:    ctx,
		cancel: cancel,
	}

	pool.GoCtx(func(ctx context.Context) {
		select {
		case <-ctx.Done():
			return
		default:
			// do something
		}
	})

	pool.Stop()

	// assert that the goroutine has stopped
}
