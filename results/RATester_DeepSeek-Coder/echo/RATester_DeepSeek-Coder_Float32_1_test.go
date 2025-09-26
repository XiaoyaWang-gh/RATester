package echo

import (
	"fmt"
	"testing"
)

func TestFloat32_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			if sourceParam == "test" {
				return "3.14"
			}
			return ""
		},
	}

	var dest float32
	b.Float32("test", &dest)

	if dest != float32(3.14) {
		t.Errorf("Expected 3.14, got %v", dest)
	}
}
