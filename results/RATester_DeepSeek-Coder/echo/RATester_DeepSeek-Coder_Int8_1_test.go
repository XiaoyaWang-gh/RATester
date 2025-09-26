package echo

import (
	"fmt"
	"testing"
)

func TestInt8_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			if sourceParam == "test" {
				return "127"
			}
			return ""
		},
	}

	var dest int8
	err := b.Int8("test", &dest).BindError()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if dest != 127 {
		t.Errorf("Expected 127, got %v", dest)
	}

	err = b.Int8("invalid", &dest).BindError()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	err = b.Int8("overflow", &dest).BindError()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
