package context

import (
	"fmt"
	"testing"
)

func TestIsNotFound_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	output := &BeegoOutput{Status: 404}
	if !output.IsNotFound() {
		t.Errorf("Expected true, got false")
	}

	output = &BeegoOutput{Status: 200}
	if output.IsNotFound() {
		t.Errorf("Expected false, got true")
	}
}
