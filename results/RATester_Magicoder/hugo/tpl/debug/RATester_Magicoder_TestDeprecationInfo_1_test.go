package debug

import (
	"fmt"
	"testing"
)

func TestTestDeprecationInfo_1(t *testing.T) {
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

	result := ns.TestDeprecationInfo(item, alternative)

	if result != "" {
		t.Errorf("Expected empty string, got %s", result)
	}
}
