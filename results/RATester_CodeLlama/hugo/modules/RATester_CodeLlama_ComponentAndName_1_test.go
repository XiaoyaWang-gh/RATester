package modules

import (
	"fmt"
	"testing"
)

func TestComponentAndName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := Mount{
		Source: "scss",
		Target: "assets/bootstrap/scss",
	}
	c, n := m.ComponentAndName()
	if c != "scss" {
		t.Errorf("expected %q, got %q", "scss", c)
	}
	if n != "scss" {
		t.Errorf("expected %q, got %q", "scss", n)
	}
}
