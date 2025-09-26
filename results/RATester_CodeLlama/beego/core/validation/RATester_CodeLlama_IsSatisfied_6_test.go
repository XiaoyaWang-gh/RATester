package validation

import (
	"fmt"
	"testing"
)

func TestIsSatisfied_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := Min{Min: 10}
	if m.IsSatisfied(10) {
		t.Errorf("IsSatisfied() = true, want false")
	}
	if !m.IsSatisfied(11) {
		t.Errorf("IsSatisfied() = false, want true")
	}
	if !m.IsSatisfied(100) {
		t.Errorf("IsSatisfied() = false, want true")
	}
	if m.IsSatisfied(9) {
		t.Errorf("IsSatisfied() = true, want false")
	}
}
