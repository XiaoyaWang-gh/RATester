package commands

import (
	"fmt"
	"testing"
)

func TestWasErr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &hugoBuilderErrState{}
	e.setWasErr(true)
	if e.wasErr() != true {
		t.Errorf("wasErr() = %v, want %v", e.wasErr(), true)
	}
}
