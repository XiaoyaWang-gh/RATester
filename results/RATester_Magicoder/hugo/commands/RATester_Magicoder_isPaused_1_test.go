package commands

import (
	"fmt"
	"testing"
)

func TestisPaused_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	state := &hugoBuilderErrState{
		paused: true,
	}

	if !state.isPaused() {
		t.Error("Expected isPaused to return true")
	}

	state.paused = false

	if state.isPaused() {
		t.Error("Expected isPaused to return false")
	}
}
