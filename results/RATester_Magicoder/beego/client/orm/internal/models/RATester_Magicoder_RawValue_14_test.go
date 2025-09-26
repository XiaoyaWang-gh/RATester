package models

import (
	"fmt"
	"testing"
)

func TestRawValue_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	field := CharField("test")
	expected := "test"
	result := field.RawValue()
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
