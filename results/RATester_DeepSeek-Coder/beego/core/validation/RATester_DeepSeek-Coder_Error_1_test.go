package validation

import (
	"fmt"
	"testing"
)

func TestError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := &Validation{}
	message := "test error"
	result := v.Error(message)
	if result.Ok != false {
		t.Errorf("Expected Ok to be false, got %v", result.Ok)
	}
	if result.Error.Message != message {
		t.Errorf("Expected Error.Message to be %s, got %s", message, result.Error.Message)
	}
	if len(v.Errors) != 1 {
		t.Errorf("Expected Errors length to be 1, got %d", len(v.Errors))
	}
	if v.Errors[0].Message != message {
		t.Errorf("Expected v.Errors[0].Message to be %s, got %s", message, v.Errors[0].Message)
	}
}
