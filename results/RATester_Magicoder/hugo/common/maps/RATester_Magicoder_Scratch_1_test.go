package maps

import (
	"fmt"
	"testing"
)

func TestScratch_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	scratch := &Scratch{}
	scratching := scratcher{s: scratch}

	result := scratching.Scratch()

	if result != scratch {
		t.Errorf("Expected %v, got %v", scratch, result)
	}
}
