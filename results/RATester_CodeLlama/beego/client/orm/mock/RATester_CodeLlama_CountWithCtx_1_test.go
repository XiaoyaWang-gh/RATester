package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestCountWithCtx_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQueryM2Mer{}
	ctx := context.Background()
	got, err := d.CountWithCtx(ctx)
	if err != nil {
		t.Fatalf("CountWithCtx() error = %v", err)
	}
	if got != 0 {
		t.Errorf("CountWithCtx() = %v, want %v", got, 0)
	}
}
