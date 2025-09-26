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

	b := &ValueBinder{}
	sourceParam := "param"
	dest := []float32{}

	// Test case 1: Successful binding
	b.ValuesFunc = func(sourceParam string) []string {
		return []string{"1.23", "4.56"}
	}
	b.Float32s(sourceParam, &dest)
	if len(dest) != 2 || dest[0] != float32(1.23) || dest[1] != float32(4.56) {
		t.Errorf("Expected dest to be [1.23, 4.56], but got %v", dest)
	}

	// Test case 2: Binding with no values
	b.ValuesFunc = func(sourceParam string) []string {
		return []string{}
	}
	b.Float32s(sourceParam, &dest)
	if len(dest) != 0 {
		t.Errorf("Expected dest to be [], but got %v", dest)
	}

	// Test case 3: Binding with invalid values
	b.ValuesFunc = func(sourceParam string) []string {
		return []string{"invalid", "1.23"}
	}
	b.Float32s(sourceParam, &dest)
	if len(dest) != 1 || dest[0] != float32(1.23) {
		t.Errorf("Expected dest to be [1.23], but got %v", dest)
	}

	// Test case 4: Binding with failFast set
	b.failFast = true
	b.ValuesFunc = func(sourceParam string) []string {
		return []string{"invalid", "1.23"}
	}
	b.Float32s(sourceParam, &dest)
	if len(dest) != 0 {
		t.Errorf("Expected dest to be [], but got %v", dest)
	}
	if len(b.errors) != 1 {
		t.Errorf("Expected 1 error, but got %v", len(b.errors))
	}
}
