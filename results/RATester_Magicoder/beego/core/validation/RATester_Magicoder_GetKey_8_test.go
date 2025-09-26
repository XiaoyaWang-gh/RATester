package validation

import (
	"fmt"
	"testing"
)

func TestGetKey_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	min := Min{
		Min: 1,
		Key: "testKey",
	}

	expected := "testKey"
	result := min.GetKey()

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
