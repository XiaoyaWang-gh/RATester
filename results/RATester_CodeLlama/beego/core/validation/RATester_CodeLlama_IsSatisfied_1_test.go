package validation

import (
	"fmt"
	"testing"
)

func TestIsSatisfied_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := Numeric{}
	obj := "1234567890"
	if !n.IsSatisfied(obj) {
		t.Errorf("IsSatisfied() = %v, want %v", false, true)
	}
}
