package codegen

import (
	"fmt"
	"reflect"
	"testing"
)

func TestinOutStr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	method := Method{
		Owner:     reflect.TypeOf(Method{}),
		OwnerName: "Method",
		Name:      "inOutStr",
		In:        []string{"arg0", "arg1"},
	}

	expected := "(arg0, arg1)"
	result := method.inOutStr()

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
