package echo

import (
	"fmt"
	"testing"
)

func TestUint8_1(t *testing.T) {
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

	var dest uint8
	err := b.Uint8("test", &dest).BindError()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if dest != 123 {
		t.Errorf("Expected 123, got %v", dest)
	}
}
