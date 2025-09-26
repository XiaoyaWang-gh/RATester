package para

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestStart_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	w := &Workers{
		sem: make(chan struct{}, 1),
	}

	r, ctx := w.Start(ctx)

	r.Run(func() error {
		time.Sleep(time.Second)
		return nil
	})

	err := r.Wait()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
