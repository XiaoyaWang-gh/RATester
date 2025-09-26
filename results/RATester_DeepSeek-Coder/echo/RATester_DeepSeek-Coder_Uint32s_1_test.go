package echo

import (
	"fmt"
	"testing"
)

func TestUint32s_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &ValueBinder{
		ValuesFunc: func(sourceParam string) []string {
			if sourceParam == "ids" {
				return []string{"1", "2", "3"}
			}
			return nil
		},
	}

	var dest []uint32
	b.Uint32s("ids", &dest)

	if len(dest) != 3 {
		t.Errorf("Expected 3 values, got %d", len(dest))
	}

	if dest[0] != 1 || dest[1] != 2 || dest[2] != 3 {
		t.Errorf("Expected [1, 2, 3], got %v", dest)
	}
}
