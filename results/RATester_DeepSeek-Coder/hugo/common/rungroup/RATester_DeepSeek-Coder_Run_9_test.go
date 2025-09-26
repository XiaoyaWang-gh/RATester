package rungroup

import (
	"context"
	"fmt"
	"testing"
)

func TestRun_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := Config[int]{
		NumWorkers: 3,
		Handle: func(ctx context.Context, v int) error {
			if v%2 == 0 {
				return fmt.Errorf("error with value: %d", v)
			}
			return nil
		},
	}

	g := Run[int](ctx, cfg)

	for i := 0; i < 10; i++ {
		if err := g.Enqueue(i); err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	}

	if err := g.Wait(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
