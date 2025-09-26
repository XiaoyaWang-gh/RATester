package echo

import (
	"fmt"
	"testing"
)

func TestInt32_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			// Implement your test case here
			return "123"
		},
	}

	var dest int32
	err := b.Int32("test", &dest)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if dest != 123 {
		t.Errorf("Expected dest to be 123, but got %v", dest)
	}
}
