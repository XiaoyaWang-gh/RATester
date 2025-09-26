package echo

import (
	"errors"
	"fmt"
	"testing"
)

func TestIntsValue_1(t *testing.T) {
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

	dest := []int{}
	b.intsValue("param", &dest, true)

	if len(b.errors) != 0 {
		t.Errorf("Expected no errors, got %v", b.errors)
	}

	if len(dest) != 3 {
		t.Errorf("Expected dest to have 3 elements, got %v", dest)
	}

	b.failFast = false
	b.intsValue("param", &dest, true)

	if len(b.errors) == 0 {
		t.Errorf("Expected errors, got none")
	}
}
