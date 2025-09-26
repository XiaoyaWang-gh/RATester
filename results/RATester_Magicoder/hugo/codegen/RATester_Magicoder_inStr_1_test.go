package codegen

import (
	"fmt"
	"reflect"
	"testing"
)

func TestinStr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	method := Method{
		Owner:     reflect.TypeOf(Method{}),
		OwnerName: "Method",
		Name:      "inStr",
		In:        []string{"arg0 string", "arg1 int"},
	}

	expected := "(arg0 string, arg1 int)"
	result := method.inStr()

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
