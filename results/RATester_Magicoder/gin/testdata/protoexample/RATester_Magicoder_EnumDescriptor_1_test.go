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
	data, path := foo.EnumDescriptor()

	if len(data) == 0 {
		t.Error("Expected data to be non-empty")
	}

	if len(path) != 1 || path[0] != 0 {
		t.Errorf("Expected path to be [0], but got %v", path)
	}
}
