package context

import (
	"fmt"
	"testing"
)

func TestIsSuccessful_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	output := &BeegoOutput{Status: 200}
	if !output.IsSuccessful() {
		t.Error("Expected true, got false")
	}

	output = &BeegoOutput{Status: 300}
	if output.IsSuccessful() {
		t.Error("Expected false, got true")
	}

	output = &BeegoOutput{Status: 199}
	if output.IsSuccessful() {
		t.Error("Expected false, got true")
	}
}
