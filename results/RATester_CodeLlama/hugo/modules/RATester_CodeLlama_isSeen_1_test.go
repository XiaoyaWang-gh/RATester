package modules

import (
	"fmt"
	"testing"
)

func TestIsSeen_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &collector{}
	c.seen = make(map[string]bool)
	if c.isSeen("") {
		t.Error("Expected false")
	}
	if !c.isSeen("") {
		t.Error("Expected true")
	}
}
