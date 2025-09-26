package protoexample

import (
	"fmt"
	"testing"
)

func TestDescriptor_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	foo := FOO(0)
	descriptor := foo.Descriptor()

	if descriptor == nil {
		t.Error("Expected a non-nil descriptor, but got nil")
	}
}
