package maps

import (
	"fmt"
	"testing"
)

func TestNewScratch_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	scratch := NewScratch()

	if scratch == nil {
		t.Errorf("Expected NewScratch() to return a non-nil value")
	}

	if scratch.values == nil {
		t.Errorf("Expected scratch.values to be initialized")
	}
}
