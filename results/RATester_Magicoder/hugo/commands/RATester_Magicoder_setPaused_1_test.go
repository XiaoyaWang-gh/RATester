package commands

import (
	"fmt"
	"testing"
)

func TestsetPaused_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	state := &hugoBuilderErrState{}

	// Test setting paused to true
	state.setPaused(true)
	if !state.isPaused() {
		t.Error("Expected paused to be true, but it was false")
	}

	// Test setting paused to false
	state.setPaused(false)
	if state.isPaused() {
		t.Error("Expected paused to be false, but it was true")
	}
}
