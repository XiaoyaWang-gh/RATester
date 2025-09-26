package validation

import (
	"fmt"
	"testing"
)

func TestTel_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := Validation{}

	// Test case 1: Valid phone number
	result := v.Tel("13800138000", "phone")
	if !result.Ok {
		t.Errorf("Expected valid phone number, but got invalid")
	}

	// Test case 2: Invalid phone number
	result = v.Tel("1234567890", "phone")
	if result.Ok {
		t.Errorf("Expected invalid phone number, but got valid")
	}

	// Test case 3: Empty phone number
	result = v.Tel("", "phone")
	if result.Ok {
		t.Errorf("Expected invalid phone number, but got valid")
	}
}
