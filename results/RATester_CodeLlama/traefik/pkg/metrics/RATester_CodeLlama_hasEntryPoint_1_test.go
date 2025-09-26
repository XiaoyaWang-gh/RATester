package metrics

import (
	"fmt"
	"testing"
)

func TestHasEntryPoint_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &dynamicConfig{
		entryPoints: map[string]bool{
			"foo": true,
		},
	}

	if !d.hasEntryPoint("foo") {
		t.Errorf("expected true")
	}

	if d.hasEntryPoint("bar") {
		t.Errorf("expected false")
	}
}
