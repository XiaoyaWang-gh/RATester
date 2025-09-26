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

	// Test initial count
	expectedCount := 0
	actualCount := init.InitCount()
	if actualCount != expectedCount {
		t.Errorf("Expected InitCount to be %d, but got %d", expectedCount, actualCount)
	}

	// Test incrementing count
	init.add(false, func(ctx context.Context) (any, error) { return nil, nil })
	expectedCount = 1
	actualCount = init.InitCount()
	if actualCount != expectedCount {
		t.Errorf("Expected InitCount to be %d, but got %d", expectedCount, actualCount)
	}
}
