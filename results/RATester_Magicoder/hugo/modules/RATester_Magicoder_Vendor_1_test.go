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

	adapter := &moduleAdapter{
		vendor: true,
	}

	if !adapter.Vendor() {
		t.Errorf("Expected Vendor() to return true, but got false")
	}
}
