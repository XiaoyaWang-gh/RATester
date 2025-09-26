package protoexample

import (
	"fmt"
	"testing"
)

func TestEnum_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	foo := FOO(123)
	result := foo.Enum()
	if *result != foo {
		t.Errorf("Expected %v, got %v", foo, *result)
	}
}
