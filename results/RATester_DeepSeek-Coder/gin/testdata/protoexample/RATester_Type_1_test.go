package protoexample

import (
	"fmt"
	"testing"
)

func TestType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	foo := FOO(0)
	expected := &file_test_proto_enumTypes[0]
	actual := foo.Type()
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
