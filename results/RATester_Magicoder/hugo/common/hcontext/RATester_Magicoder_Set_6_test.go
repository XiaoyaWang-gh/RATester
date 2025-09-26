package hcontext

import (
	"context"
	"fmt"
	"testing"
)

func TestSet_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	key := keyInContext[int, string]{id: "test"}
	ctx = key.Set(ctx, 123)
	value := ctx.Value(key.id)
	if value != 123 {
		t.Errorf("Expected value to be 123, but got %v", value)
	}
}
