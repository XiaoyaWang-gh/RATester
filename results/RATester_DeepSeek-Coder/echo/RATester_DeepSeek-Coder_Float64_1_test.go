package echo

import (
	"fmt"
	"testing"
)

func TestFloat64_1(t *testing.T) {
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

	var dest float64
	b.Float64("test", &dest)

	if dest != 1.23 {
		t.Errorf("Expected 1.23, got %v", dest)
	}
}
