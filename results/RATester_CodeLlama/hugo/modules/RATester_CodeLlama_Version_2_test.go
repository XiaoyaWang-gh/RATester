package modules

import (
	"fmt"
	"testing"
)

func TestVersion_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &moduleAdapter{}
	m.version = "1.0.0"
	if m.Version() != "1.0.0" {
		t.Errorf("Version() = %v, want %v", m.Version(), "1.0.0")
	}
}
