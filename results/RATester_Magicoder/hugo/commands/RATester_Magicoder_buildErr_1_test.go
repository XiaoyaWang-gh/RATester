package commands

import (
	"errors"
	"fmt"
	"sync"
	"testing"
)

func TestbuildErr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	state := &hugoBuilderErrState{
		mu:       sync.Mutex{},
		paused:   false,
		builderr: nil,
		waserr:   false,
	}

	err := state.buildErr()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	state.setBuildErr(errors.New("test error"))
	err = state.buildErr()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
