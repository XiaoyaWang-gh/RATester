package codegen

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeclaration_1(t *testing.T) {
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

	expected := "func (m interface{}) TestMethod(string, int) error"
	result := method.Declaration("m")

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
