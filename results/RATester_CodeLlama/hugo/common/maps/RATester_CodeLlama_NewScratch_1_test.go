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
		t.Errorf("NewScratch() returned nil")
	}
	if scratch.values == nil {
		t.Errorf("NewScratch() returned a Scratch with nil values")
	}
}
