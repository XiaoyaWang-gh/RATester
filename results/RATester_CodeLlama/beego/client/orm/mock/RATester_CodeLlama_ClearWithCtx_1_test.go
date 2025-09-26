package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestClearWithCtx_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQueryM2Mer{}
	ctx := context.Background()
	got, err := d.ClearWithCtx(ctx)
	if err != nil {
		t.Fatalf("ClearWithCtx() error = %v", err)
	}
	if got != 0 {
		t.Errorf("ClearWithCtx() got = %v, want %v", got, 0)
	}
}
