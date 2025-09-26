package web

import (
	"fmt"
	"testing"
)

func TestAddAPPStartHook_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	hooks = []hookfunc{}

	AddAPPStartHook(func() error {
		return nil
	})

	if len(hooks) != 1 {
		t.Errorf("Expected 1 hook, got %d", len(hooks))
	}

	AddAPPStartHook(func() error {
		return nil
	})

	if len(hooks) != 2 {
		t.Errorf("Expected 2 hooks, got %d", len(hooks))
	}
}
