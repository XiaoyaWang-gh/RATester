package debug

import (
	"fmt"
	"sync"
	"testing"
)

func TestTestDeprecationErr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{
		timersMu: sync.Mutex{},
		timers:   make(map[string][]*timer),
	}

	item := "item"
	alternative := "alternative"

	result := ns.TestDeprecationErr(item, alternative)

	if result != "" {
		t.Errorf("Expected empty string, got %s", result)
	}
}
