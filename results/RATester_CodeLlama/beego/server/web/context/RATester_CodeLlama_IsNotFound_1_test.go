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
		t.Errorf("IsNotFound() = %v, want %v", output.IsNotFound(), true)
	}
}
