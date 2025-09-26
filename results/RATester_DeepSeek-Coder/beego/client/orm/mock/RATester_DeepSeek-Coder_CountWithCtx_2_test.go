package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestCountWithCtx_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	ctx := context.Background()
	d := &DoNothingQuerySetter{}

	count, err := d.CountWithCtx(ctx)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if count != 0 {
		t.Errorf("Expected count to be 0, got %v", count)
	}
}
