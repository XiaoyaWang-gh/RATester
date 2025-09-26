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
		t.Error("Expected a new Scratch, but got nil")
	}

	if scratch.values == nil {
		t.Error("Expected a new map, but got nil")
	}
}
