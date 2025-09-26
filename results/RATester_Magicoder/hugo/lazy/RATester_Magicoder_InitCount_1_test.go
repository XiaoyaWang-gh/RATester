package lazy

import (
	"context"
	"fmt"
	"testing"
)

func TestInitCount_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	init := &Init{}

	// Test case 1: InitCount should return the correct count
	expectedCount := 5
	for i := 0; i < expectedCount; i++ {
		init.Add(func(ctx context.Context) (any, error) { return nil, nil })
	}
	if init.InitCount() != expectedCount {
		t.Errorf("Expected InitCount to return %d, but got %d", expectedCount, init.InitCount())
	}

	// Test case 2: InitCount should return 0 when no dependencies are added
	init.Reset()
	if init.InitCount() != 0 {
		t.Errorf("Expected InitCount to return 0, but got %d", init.InitCount())
	}
}
