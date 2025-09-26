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

	ctx := context.Background()
	dnm := &DoNothingQueryM2Mer{}
	count, err := dnm.CountWithCtx(ctx)
	if err != nil {
		t.Errorf("CountWithCtx returned an error: %v", err)
	}
	if count != 0 {
		t.Errorf("CountWithCtx returned %d, expected 0", count)
	}
}
