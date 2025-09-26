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
	if got, want := foo.Descriptor(), file_test_proto_enumTypes[0].Descriptor(); got != want {
		t.Errorf("Descriptor() = %v, want %v", got, want)
	}
}
