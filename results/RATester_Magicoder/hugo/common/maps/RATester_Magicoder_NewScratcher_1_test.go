package maps

import (
	"fmt"
	"testing"
)

func TestNewScratcher_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	scratch := NewScratcher()
	if scratch == nil {
		t.Error("Expected a non-nil Scratcher, got nil")
	}
}
