package models

import (
	"fmt"
	"testing"
)

func TestRawValue_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	field := PositiveIntegerField(10)
	expected := uint32(10)

	result := field.RawValue()

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
