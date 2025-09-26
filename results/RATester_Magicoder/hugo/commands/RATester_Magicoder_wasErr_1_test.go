package commands

import (
	"fmt"
	"testing"
)

func TestwasErr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	state := &hugoBuilderErrState{}

	// Test case 1: wasErr is false
	state.setWasErr(false)
	if state.wasErr() != false {
		t.Errorf("Expected wasErr to be false, but got true")
	}

	// Test case 2: wasErr is true
	state.setWasErr(true)
	if state.wasErr() != true {
		t.Errorf("Expected wasErr to be true, but got false")
	}
}
