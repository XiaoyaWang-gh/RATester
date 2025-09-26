package echo

import (
	"fmt"
	"testing"
)

func TestMustInt16s_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			if sourceParam == "test" {
				return "1"
			}
			return ""
		},
	}

	var dest []int16
	err := b.MustInt16s("test", &dest).BindError()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(dest) != 1 {
		t.Errorf("Expected 1 element, got %d", len(dest))
	}
	if dest[0] != 1 {
		t.Errorf("Expected 1, got %d", dest[0])
	}

	err = b.MustInt16s("missing", &dest).BindError()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
