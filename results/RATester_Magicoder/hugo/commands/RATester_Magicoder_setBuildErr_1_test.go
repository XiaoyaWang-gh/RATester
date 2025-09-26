package commands

import (
	"errors"
	"fmt"
	"testing"
)

func TestsetBuildErr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	state := &hugoBuilderErrState{}
	err := errors.New("test error")
	state.setBuildErr(err)
	if state.builderr != err {
		t.Errorf("Expected %v, got %v", err, state.builderr)
	}
}
