package mirror

import (
	"fmt"
	"testing"
)

func TestAddMirror_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mirroring := &Mirroring{}

	// Test case 1: Percent is less than 0
	err := mirroring.AddMirror(nil, -1)
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	// Test case 2: Percent is greater than 100
	err = mirroring.AddMirror(nil, 101)
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	// Test case 3: Percent is between 0 and 100
	err = mirroring.AddMirror(nil, 50)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
