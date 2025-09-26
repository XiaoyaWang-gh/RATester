package orm

import (
	"fmt"
	"testing"
)

func TestRegisterDriver_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	drivers := make(map[string]DriverType)

	// Test case 1: Register a new driver
	err := RegisterDriver("newDriver", 1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if _, ok := drivers["newDriver"]; !ok {
		t.Errorf("Driver not registered")
	}

	// Test case 2: Register an existing driver of different type
	err = RegisterDriver("newDriver", 2)
	if err == nil {
		t.Errorf("Expected error not received")
	}

	// Test case 3: Register an existing driver of same type
	err = RegisterDriver("newDriver", 1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
