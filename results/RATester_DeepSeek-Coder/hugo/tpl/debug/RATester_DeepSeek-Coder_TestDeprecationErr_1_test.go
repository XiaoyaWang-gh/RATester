package debug

import (
	"fmt"
	"testing"
)

func TestTestDeprecationErr_1(t *testing.T) {
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

	expected := ""
	actual := ns.TestDeprecationErr(item, alternative)

	if actual != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}
