package echo

import (
	"fmt"
	"testing"
)

func TestInts_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			return "1,2,3"
		},
	}

	var dest []int
	b.Ints("sourceParam", &dest)

	if len(dest) != 3 {
		t.Errorf("Expected 3 elements, got %d", len(dest))
	}

	if dest[0] != 1 || dest[1] != 2 || dest[2] != 3 {
		t.Errorf("Expected [1,2,3], got %v", dest)
	}
}
