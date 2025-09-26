package hugolib

import (
	"fmt"
	"testing"
)

func TestIsDefaultCodeBlockRenderer_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	hr := hookRendererTemplate{}
	result := hr.IsDefaultCodeBlockRenderer()
	if result != false {
		t.Errorf("Expected false, but got %v", result)
	}
}
