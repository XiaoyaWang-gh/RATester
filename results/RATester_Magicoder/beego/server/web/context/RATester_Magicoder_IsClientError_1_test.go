package context

import (
	"fmt"
	"testing"
)

func TestIsClientError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	output := &BeegoOutput{
		Status: 400,
	}

	if !output.IsClientError() {
		t.Error("Expected true, got false")
	}

	output.Status = 399
	if output.IsClientError() {
		t.Error("Expected false, got true")
	}

	output.Status = 500
	if output.IsClientError() {
		t.Error("Expected false, got true")
	}

	output.Status = 501
	if !output.IsClientError() {
		t.Error("Expected true, got false")
	}
}
