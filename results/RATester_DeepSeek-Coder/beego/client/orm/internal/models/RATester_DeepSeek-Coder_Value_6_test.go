package models

import (
	"fmt"
	"testing"
)

func TestValue_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := PositiveSmallIntegerField(65535)
	expected := uint16(65535)
	result := e.Value()
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
