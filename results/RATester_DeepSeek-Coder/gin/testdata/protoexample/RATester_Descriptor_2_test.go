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
	desc := foo.Descriptor()
	if desc == nil {
		t.Errorf("Expected a non-nil descriptor, got nil")
	}
}
