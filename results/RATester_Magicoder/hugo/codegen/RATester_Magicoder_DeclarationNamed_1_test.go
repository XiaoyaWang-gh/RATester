package codegen

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeclarationNamed_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	method := Method{
		Owner:     reflect.TypeOf(new(interface{})),
		OwnerName: "interface{}",
		Name:      "TestMethod",
		In:        []string{"string", "int"},
		Out:       []string{"error"},
	}

	expected := "func (t interface{}) TestMethod(string, int) error"
	result := method.DeclarationNamed("t")

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
