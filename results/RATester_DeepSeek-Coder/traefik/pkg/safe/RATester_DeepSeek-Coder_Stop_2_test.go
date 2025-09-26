package safe

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestStop_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

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
			time.Sleep(time.Second)
		}
	})

	pool.Stop()

	select {
	case <-ctx.Done():
		if ctx.Err() != context.Canceled {
			t.Errorf("expected context error to be context.Canceled, got %v", ctx.Err())
		}
	default:
		t.Errorf("expected context to be done")
	}
}
