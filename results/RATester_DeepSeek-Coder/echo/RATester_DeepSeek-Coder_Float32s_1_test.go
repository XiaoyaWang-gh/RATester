package echo

import (
	"fmt"
	"testing"
)

func TestFloat32s_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			if sourceParam == "test" {
				return "1.23"
			}
			return ""
		},
	}

	var dest []float32
	err := b.Float32s("test", &dest)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(dest) != 1 {
		t.Errorf("Expected 1 element, got %d", len(dest))
	}

	if dest[0] != float32(1.23) {
		t.Errorf("Expected 1.23, got %f", dest[0])
	}
}
