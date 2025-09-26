package models

import (
	"fmt"
	"testing"
)

func TestValue_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := PositiveBigIntegerField(18446744073709551615)
	expected := uint64(18446744073709551615)
	result := e.Value()
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
