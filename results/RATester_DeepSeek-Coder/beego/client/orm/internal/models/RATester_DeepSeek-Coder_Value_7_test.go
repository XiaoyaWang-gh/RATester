package models

import (
	"fmt"
	"testing"
)

func TestValue_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := SmallIntegerField(0)
	expected := int16(0)
	result := e.Value()
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	e = SmallIntegerField(-32768)
	expected = int16(-32768)
	result = e.Value()
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	e = SmallIntegerField(32767)
	expected = int16(32767)
	result = e.Value()
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
