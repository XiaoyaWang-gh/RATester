package validation

import (
	"fmt"
	"testing"
)

func TestGetKey_17(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	maxSize := MaxSize{
		Max: 10,
		Key: "testKey",
	}

	expected := "testKey"
	result := maxSize.GetKey()

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
