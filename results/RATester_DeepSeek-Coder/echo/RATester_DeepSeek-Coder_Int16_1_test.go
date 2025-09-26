package echo

import (
	"fmt"
	"testing"
)

func TestInt16_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			if sourceParam == "test" {
				return "12345"
			}
			return ""
		},
	}

	var dest int16
	err := b.Int16("test", &dest).BindError()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if dest != 12345 {
		t.Errorf("Expected 12345, got %v", dest)
	}

	err = b.Int16("invalid", &dest).BindError()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
