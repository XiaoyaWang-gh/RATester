package commands

import (
	"fmt"
	"testing"
)

func TestSetPaused_1(t *testing.T) {
	state := &hugoBuilderErrState{}

	t.Run("setPaused", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		state.setPaused(true)
		if state.paused != true {
			t.Errorf("Expected paused to be true, got %v", state.paused)
		}
	})

	t.Run("setPaused", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		state.setPaused(false)
		if state.paused != false {
			t.Errorf("Expected paused to be false, got %v", state.paused)
		}
	})
}
