package validation

import (
	"fmt"
	"testing"
)

func TestGetKey_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	max := Max{
		Max: 10,
		Key: "testKey",
	}

	expected := "testKey"
	result := max.GetKey()

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
