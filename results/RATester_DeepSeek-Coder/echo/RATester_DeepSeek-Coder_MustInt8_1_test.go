package echo

import (
	"fmt"
	"testing"
)

func TestMustInt8_1(t *testing.T) {
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

	var i int8
	err := b.MustInt8("test", &i).BindError()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if i != 127 {
		t.Errorf("Expected 127, got %v", i)
	}

	err = b.MustInt8("missing", &i).BindError()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
