package mock

import (
	"context"
	"fmt"
	"testing"
)

func TestExistWithCtx_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	d := &DoNothingQueryM2Mer{}

	result := d.ExistWithCtx(ctx, "test")

	if result != true {
		t.Errorf("Expected true, got %v", result)
	}
}
