package modules

import (
	"fmt"
	"testing"
)

func TestComponent_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := Mount{
		Source: "scss",
		Target: "assets/bootstrap/scss",
	}
	if got := m.Component(); got != "scss" {
		t.Errorf("Mount.Component() = %v, want %v", got, "scss")
	}
}
