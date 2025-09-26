package echo

import (
	"fmt"
	"testing"
)

func TestFloat_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			return "123.45"
		},
		ErrorFunc: func(sourceParam string, values []string, message interface{}, internalError error) error {
			return fmt.Errorf("error: %v", message)
		},
	}

	var dest float64
	err := b.float("test", "123.45", &dest, 64)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if dest != 123.45 {
		t.Errorf("Expected dest to be 123.45, but got %v", dest)
	}

	b.ValueFunc = func(sourceParam string) string {
		return "invalid"
	}
	err = b.float("test", "invalid", &dest, 64)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
}
