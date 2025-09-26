package commands

import (
	"fmt"
	"testing"
)

func TestsetWasErr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	state := &hugoBuilderErrState{}
	state.setWasErr(true)
	if !state.wasErr() {
		t.Error("Expected wasErr to be true, but it was false")
	}
}
