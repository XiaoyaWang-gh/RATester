package echo

import (
	"fmt"
	"testing"
)

func TestInt32_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			if sourceParam == "test" {
				return "123"
			}
			return ""
		},
	}

	var dest int32
	err := b.Int32("test", &dest).BindError()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if dest != 123 {
		t.Errorf("Expected 123, got %v", dest)
	}

	err = b.Int32("invalid", &dest).BindError()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
