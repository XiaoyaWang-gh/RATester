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
		t.Error("Expected IsForbidden to return true, but it didn't")
	}

	output = &BeegoOutput{Status: 404}
	if output.IsForbidden() {
		t.Error("Expected IsForbidden to return false, but it didn't")
	}
}
