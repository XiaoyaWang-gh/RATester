package context

import (
	"fmt"
	"testing"
)

func TestIsServerError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	output := &BeegoOutput{Status: 500}
	if !output.IsServerError() {
		t.Error("Expected true, got", output.IsServerError())
	}

	output = &BeegoOutput{Status: 499}
	if output.IsServerError() {
		t.Error("Expected false, got", output.IsServerError())
	}

	output = &BeegoOutput{Status: 600}
	if output.IsServerError() {
		t.Error("Expected false, got", output.IsServerError())
	}
}
