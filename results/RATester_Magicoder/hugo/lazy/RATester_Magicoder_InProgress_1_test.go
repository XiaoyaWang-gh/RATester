package lazy

import (
	"fmt"
	"testing"
)

func TestInProgress_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &onceMore{}

	// Test when the lock is not in progress
	if o.InProgress() {
		t.Error("Expected InProgress() to return false when the lock is not in progress")
	}

	// Test when the lock is in progress
	o.lock = 1
	if !o.InProgress() {
		t.Error("Expected InProgress() to return true when the lock is in progress")
	}

	// Test when the lock is done
	o.lock = 0
	if o.InProgress() {
		t.Error("Expected InProgress() to return false when the lock is done")
	}
}
