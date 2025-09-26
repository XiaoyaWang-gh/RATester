package validation

import (
	"fmt"
	"testing"
)

func TestIsSatisfied_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := Length{N: 10}
	obj := "hello"
	if !l.IsSatisfied(obj) {
		t.Errorf("IsSatisfied() = %v, want %v", false, true)
	}
}
