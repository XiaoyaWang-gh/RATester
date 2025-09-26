package modules

import (
	"fmt"
	"testing"
)

func TestPath_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &moduleAdapter{}
	m.path = "path"
	if m.Path() != "path" {
		t.Errorf("expected %s, got %s", "path", m.Path())
	}
	m.gomod = &goModule{}
	m.gomod.Path = "gomodpath"
	if m.Path() != "gomodpath" {
		t.Errorf("expected %s, got %s", "gomodpath", m.Path())
	}
}
