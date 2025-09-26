package echo

import (
	"errors"
	"fmt"
	"testing"
)

func TestUintsValue_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &ValueBinder{
		ValuesFunc: func(sourceParam string) []string {
			return []string{"1", "2", "3"}
		},
		ErrorFunc: func(sourceParam string, values []string, message interface{}, internalError error) error {
			return errors.New("error")
		},
		failFast: true,
	}

	var dest []uint
	b.uintsValue("sourceParam", &dest, true)

	if len(b.errors) != 0 {
		t.Errorf("Expected no errors, got %v", b.errors)
	}

	if len(dest) != 3 {
		t.Errorf("Expected dest to have 3 elements, got %d", len(dest))
	}

	if dest[0] != 1 || dest[1] != 2 || dest[2] != 3 {
		t.Errorf("Expected dest to be [1, 2, 3], got %v", dest)
	}
}
