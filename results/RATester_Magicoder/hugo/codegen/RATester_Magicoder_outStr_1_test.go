package codegen

import (
	"fmt"
	"testing"
)

func TestoutStr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	method := Method{
		Out: []string{"string", "int"},
	}

	expected := "(string, int)"
	result := method.outStr()

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
