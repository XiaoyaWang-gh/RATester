package commands

import (
	"fmt"
	"sync"
	"testing"
)

func TestIsPaused_1(t *testing.T) {
	state := &hugoBuilderErrState{
		mu:     sync.Mutex{},
		paused: false,
	}

	t.Run("isPaused", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		state.setPaused(true)
		if !state.isPaused() {
			t.Errorf("Expected isPaused to return true, got false")
		}
	})
}
