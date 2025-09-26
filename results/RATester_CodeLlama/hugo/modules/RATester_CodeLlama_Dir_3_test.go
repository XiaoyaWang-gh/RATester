package modules

import (
	"fmt"
	"testing"
)

func TestDir_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &moduleAdapter{}
	m.gomod = &goModule{}
	m.gomod.Dir = "test"
	if m.Dir() != "test" {
		t.Errorf("Dir() = %v, want %v", m.Dir(), "test")
	}
}
