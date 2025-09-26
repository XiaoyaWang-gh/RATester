package validation

import (
	"fmt"
	"testing"
)

func TestIsSatisfied_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := Phone{}
	obj := "13800138000"
	if !p.IsSatisfied(obj) {
		t.Errorf("Phone.IsSatisfied() = %v, want %v", false, true)
	}
}
