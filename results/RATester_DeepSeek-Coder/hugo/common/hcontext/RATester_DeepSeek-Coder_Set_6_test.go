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

	type testKey struct{}
	key := keyInContext[int, testKey]{id: testKey{}}
	ctx := context.Background()
	newCtx := key.Set(ctx, 1)
	val := newCtx.Value(key.id)
	if val != 1 {
		t.Errorf("Expected 1, got %v", val)
	}
}
