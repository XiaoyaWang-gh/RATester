package models

import (
	"fmt"
	"testing"
)

func TestRawValue_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := BigIntegerField(1234567890)
	expected := int64(1234567890)
	result := e.RawValue()
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
