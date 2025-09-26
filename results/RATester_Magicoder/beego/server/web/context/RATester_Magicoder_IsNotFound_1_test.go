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
		t.Error("Expected IsNotFound to return true, but it didn't")
	}

	output = &BeegoOutput{Status: 200}
	if output.IsNotFound() {
		t.Error("Expected IsNotFound to return false, but it didn't")
	}
}
