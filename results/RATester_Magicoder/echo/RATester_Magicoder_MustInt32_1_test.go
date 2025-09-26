package echo

import (
	"fmt"
	"testing"
)

func TestMustInt32_1(t *testing.T) {
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
	err := b.MustInt32("test", &dest).BindError()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if dest != 123 {
		t.Errorf("Expected dest to be 123, got %v", dest)
	}

	err = b.MustInt32("invalid", &dest).BindError()
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
