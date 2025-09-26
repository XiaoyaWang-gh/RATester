package cache

import (
	"context"
	"fmt"
	"testing"
)

func TestClearAll_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fc := &FileCache{
		CachePath: "/tmp/test",
	}

	ctx := context.Background()
	err := fc.ClearAll(ctx)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
