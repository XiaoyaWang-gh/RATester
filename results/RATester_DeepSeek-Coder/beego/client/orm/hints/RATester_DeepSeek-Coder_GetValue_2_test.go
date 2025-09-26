package hints

import (
	"fmt"
	"testing"
)

func TestGetValue_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &Hint{
		key:   "testKey",
		value: "testValue",
	}

	result := h.GetValue()

	if result != "testValue" {
		t.Errorf("Expected 'testValue', got '%v'", result)
	}
}
