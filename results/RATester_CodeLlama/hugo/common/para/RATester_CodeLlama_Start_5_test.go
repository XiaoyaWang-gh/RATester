package para

import (
	"context"
	"fmt"
	"testing"
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
	w.sem <- struct{}{}
	r, _ := w.Start(ctx)
	if r == nil {
		t.Error("expected a non-nil Runner")
	}
	if w.sem == nil {
		t.Error("expected a non-nil sem")
	}
	if len(w.sem) != 0 {
		t.Errorf("expected sem to be empty, got %d", len(w.sem))
	}
}
