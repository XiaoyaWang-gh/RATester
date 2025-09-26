package debug

import (
	"fmt"
	"testing"
)

func TestTestDeprecationWarn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{
		timers: make(map[string][]*timer),
	}

	item := "item"
	alternative := "alternative"

	result := ns.TestDeprecationWarn(item, alternative)

	if result != "" {
		t.Errorf("Expected an empty string, but got: %s", result)
	}
}
