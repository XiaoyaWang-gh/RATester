package echo

import (
	"fmt"
	"testing"
)

func TestMustInt64_1(t *testing.T) {
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

	var dest int64
	err := b.MustInt64("test", &dest).BindError()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if dest != 123 {
		t.Errorf("Expected dest to be 123, but got %v", dest)
	}

	err = b.MustInt64("invalid", &dest).BindError()
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}
