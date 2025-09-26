package context

import (
	"fmt"
	"testing"
)

func TestIsForbidden_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	output := &BeegoOutput{Status: 403}
	if !output.IsForbidden() {
		t.Errorf("Expected true, got false")
	}

	output = &BeegoOutput{Status: 404}
	if output.IsForbidden() {
		t.Errorf("Expected false, got true")
	}
}
