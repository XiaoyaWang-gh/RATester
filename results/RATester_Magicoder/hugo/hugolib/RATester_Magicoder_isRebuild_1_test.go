package hugolib

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func TestisRebuild_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &HugoSites{
		buildCounter: atomic.Uint64{},
	}

	// Test case 1: buildCounter is greater than 0
	h.buildCounter.Store(1)
	if !h.isRebuild() {
		t.Error("Expected isRebuild to return true, but it returned false")
	}

	// Test case 2: buildCounter is 0
	h.buildCounter.Store(0)
	if h.isRebuild() {
		t.Error("Expected isRebuild to return false, but it returned true")
	}
}
