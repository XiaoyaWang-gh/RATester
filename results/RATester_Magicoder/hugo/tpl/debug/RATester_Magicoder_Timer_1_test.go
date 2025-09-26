package debug

import (
	"fmt"
	"testing"
)

func TestTimer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{
		timers: make(map[string][]*timer),
	}

	timer := ns.Timer("test")
	if timer == nil {
		t.Error("Expected a timer, got nil")
	}

	if len(ns.timers) != 1 {
		t.Errorf("Expected 1 timer, got %d", len(ns.timers))
	}

	if _, ok := ns.timers["test"]; !ok {
		t.Error("Expected timer for 'test'")
	}
}
