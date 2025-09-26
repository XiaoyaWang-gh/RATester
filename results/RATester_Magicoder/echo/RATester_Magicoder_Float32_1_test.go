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
			// Implement your test case here
			return "1.23"
		},
	}

	var dest float32
	err := b.Float32("test", &dest)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if dest != 1.23 {
		t.Errorf("Expected dest to be 1.23, but got %v", dest)
	}
}
