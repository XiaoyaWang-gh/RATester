package codegen

import (
	"fmt"
	"reflect"
	"testing"
)

func TestString_18(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	method := Method{
		Owner:     reflect.TypeOf(new(interface{})),
		OwnerName: "interface{}",
		Name:      "TestMethod",
		Imports:   []string{"fmt"},
		In:        []string{"string", "int"},
		Out:       []string{"string", "error"},
	}

	expected := "TestMethod(string, int) (string, error)\n"
	result := method.String()

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
