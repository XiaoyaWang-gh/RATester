package hints

import (
	"fmt"
	"testing"
)

func TestGetKey_23(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &Hint{
		key:   "testKey",
		value: "testValue",
	}

	expected := "testKey"
	result := h.GetKey()

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
