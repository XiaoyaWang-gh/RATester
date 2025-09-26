package echo

import (
	"fmt"
	"testing"
)

func TestMustUint8s_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			if sourceParam == "test" {
				return "123"
			}
			return ""
		},
	}

	var dest []uint8
	err := b.MustUint8s("test", &dest).BindError()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(dest) != 1 || dest[0] != 123 {
		t.Errorf("Expected [123], got %v", dest)
	}

	err = b.MustUint8s("missing", &dest).BindError()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
