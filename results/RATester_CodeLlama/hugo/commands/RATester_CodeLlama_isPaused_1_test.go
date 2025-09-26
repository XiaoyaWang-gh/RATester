package commands

import (
	"fmt"
	"testing"
)

func TestIsPaused_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &hugoBuilderErrState{}
	if e.isPaused() {
		t.Errorf("isPaused() = true, want false")
	}
	e.setPaused(true)
	if !e.isPaused() {
		t.Errorf("isPaused() = false, want true")
	}
}
