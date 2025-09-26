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

	err := fc.ClearAll(context.Background())
	if err != nil {
		t.Errorf("ClearAll failed: %v", err)
	}
}
