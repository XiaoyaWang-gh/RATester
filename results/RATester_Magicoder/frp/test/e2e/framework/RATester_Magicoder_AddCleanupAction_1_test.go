package framework

import (
	"fmt"
	"testing"
)

func TestAddCleanupAction_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fn := func() {
		// Your cleanup function here
	}

	handle := AddCleanupAction(fn)

	if handle == nil {
		t.Fatal("Expected a non-nil handle, but got nil")
	}

	// Add more test cases as needed
}
