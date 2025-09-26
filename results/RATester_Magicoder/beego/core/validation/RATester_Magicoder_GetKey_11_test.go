package validation

import (
	"fmt"
	"testing"
)

func TestGetKey_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testTel := Tel{
		Key: "testKey",
	}

	result := testTel.GetKey()

	if result != "testKey" {
		t.Errorf("Expected 'testKey', but got '%s'", result)
	}
}
