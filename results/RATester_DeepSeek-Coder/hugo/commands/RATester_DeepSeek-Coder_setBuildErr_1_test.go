package commands

import (
	"errors"
	"fmt"
	"testing"
)

func TestSetBuildErr_1(t *testing.T) {
	state := &hugoBuilderErrState{}

	t.Run("setBuildErr", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		state.setBuildErr(errors.New("test error"))
		if state.builderr != errors.New("test error") {
			t.Errorf("Expected error to be 'test error', got %v", state.builderr)
		}
	})
}
