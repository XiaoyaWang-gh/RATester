package metrics

import (
	"fmt"
	"testing"
)

func TestHasRouter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &dynamicConfig{
		routers: map[string]bool{
			"router1": true,
			"router2": true,
		},
	}
	if !d.hasRouter("router1") {
		t.Errorf("expected true, got false")
	}
	if d.hasRouter("router3") {
		t.Errorf("expected false, got true")
	}
}
