package hints

import (
	"fmt"
	"testing"
)

func TestNewHint_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	key := "testKey"
	value := "testValue"
	hint := NewHint(key, value)

	if hint.GetKey() != key {
		t.Errorf("Expected key to be %v, but got %v", key, hint.GetKey())
	}

	if hint.GetValue() != value {
		t.Errorf("Expected value to be %v, but got %v", value, hint.GetValue())
	}
}
