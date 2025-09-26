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

	var (
		ctx = context.Background()
		cfg = Config[int]{
			NumWorkers: 2,
			Handle: func(ctx context.Context, v int) error {
				return nil
			},
		}
	)

	g := Run[int](ctx, cfg)

	for i := 0; i < 10; i++ {
		g.Enqueue(i)
	}

	if err := g.Wait(); err != nil {
		t.Fatal(err)
	}
}
