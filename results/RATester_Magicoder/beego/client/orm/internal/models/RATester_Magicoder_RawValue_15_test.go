package models

import (
	"fmt"
	"testing"
)

func TestRawValue_15(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	field := PositiveBigIntegerField(1234567890)
	expected := field.Value()
	actual := field.RawValue()
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
