package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestPrepareInsertWithCtx_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	d := &DoNothingQuerySetter{}
	_, err := d.PrepareInsertWithCtx(ctx)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
