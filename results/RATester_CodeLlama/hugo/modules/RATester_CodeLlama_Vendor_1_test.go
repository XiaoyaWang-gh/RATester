package modules

import (
	"fmt"
	"testing"
)

func TestVendor_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &moduleAdapter{}
	if m.Vendor() {
		t.Errorf("Vendor() = %v, want %v", m.Vendor(), false)
	}
}
