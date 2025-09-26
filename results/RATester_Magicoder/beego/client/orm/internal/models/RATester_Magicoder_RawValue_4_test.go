package models

import (
	"fmt"
	"testing"
)

func TestRawValue_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	field := PositiveSmallIntegerField(10)
	expected := uint16(10)

	result := field.RawValue()

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
