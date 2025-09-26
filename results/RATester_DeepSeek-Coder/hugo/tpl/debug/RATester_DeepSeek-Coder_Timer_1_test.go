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
		t.Errorf("Expected a Timer, got nil")
	}

	ns.timersMu.Lock()
	timers := ns.timers["test"]
	ns.timersMu.Unlock()

	if len(timers) != 1 {
		t.Errorf("Expected 1 timer, got %d", len(timers))
	}

	if timers[0].start.IsZero() {
		t.Errorf("Expected start time to be set")
	}
}
