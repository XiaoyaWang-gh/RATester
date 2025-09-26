package protoexample

import (
	"fmt"
	"testing"
)

func TestEnumDescriptor_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	foo := FOO(0)
	desc, path := foo.EnumDescriptor()

	if desc == nil {
		t.Errorf("Expected a non-nil descriptor, got nil")
	}

	if len(path) != 1 || path[0] != 0 {
		t.Errorf("Expected path [0], got %v", path)
	}
}
