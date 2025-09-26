package internal

import (
	"fmt"
	"testing"
)

func TestValue_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	k := ResourceTransformationKey{
		Name: "test",
	}

	if k.Value() != "test" {
		t.Errorf("expected %q, got %q", "test", k.Value())
	}

	k.elements = []any{"test"}

	if k.Value() != "test_test" {
		t.Errorf("expected %q, got %q", "test_test", k.Value())
	}
}
