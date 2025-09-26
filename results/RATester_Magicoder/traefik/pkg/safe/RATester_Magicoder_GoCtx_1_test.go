package safe

import (
	"context"
	"fmt"
	"testing"
)

func TestGoCtx_1(t *testing.T) {
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

	routine := func(ctx context.Context) {
		select {
		case <-ctx.Done():
			return
		default:
			// do something
		}
	}

	pool.GoCtx(routine)

	// wait for the goroutine to finish
	pool.waitGroup.Wait()
}
