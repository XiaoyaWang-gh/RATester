package protoexample

import (
	"fmt"
	"testing"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
)

func TestNumber_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	foo := FOO(1)
	expected := protoreflect.EnumNumber(1)
	result := foo.Number()
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
