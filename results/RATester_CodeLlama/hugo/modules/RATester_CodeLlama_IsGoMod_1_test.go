package modules

import (
	"fmt"
	"testing"
)

func TestIsGoMod_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &moduleAdapter{}
	if m.IsGoMod() {
		t.Errorf("IsGoMod() = true, want false")
	}
}
