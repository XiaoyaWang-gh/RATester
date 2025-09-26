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

	hint := Hint{
		key:   "testKey",
		value: "testValue",
	}

	result := hint.GetValue()

	if result != "testValue" {
		t.Errorf("Expected 'testValue', but got '%v'", result)
	}
}
