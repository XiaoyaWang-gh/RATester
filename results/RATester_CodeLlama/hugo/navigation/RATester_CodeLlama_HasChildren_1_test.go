package navigation

import (
	"fmt"
	"testing"
)

func TestHasChildren_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &MenuEntry{}
	if m.HasChildren() {
		t.Errorf("Expected false, got true")
	}
	m.Children = make(Menu, 0)
	if !m.HasChildren() {
		t.Errorf("Expected true, got false")
	}
}
