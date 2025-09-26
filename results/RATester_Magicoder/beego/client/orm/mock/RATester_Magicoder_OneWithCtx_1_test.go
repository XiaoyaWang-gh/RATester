package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestOneWithCtx_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	d := &DoNothingQuerySetter{}

	err := d.OneWithCtx(ctx, nil)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
