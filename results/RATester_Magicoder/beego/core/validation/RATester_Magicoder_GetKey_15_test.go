package validation

import (
	"fmt"
	"testing"
)

func TestGetKey_15(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	alphaDash := AlphaDash{
		Key: "testKey",
	}

	result := alphaDash.GetKey()

	if result != "testKey" {
		t.Errorf("Expected 'testKey', but got '%s'", result)
	}
}
