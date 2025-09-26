package codegen

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDelegate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	method := Method{
		Owner:     reflect.TypeOf(new(interface{})),
		OwnerName: "interface{}",
		Name:      "MethodName",
		In:        []string{"string", "int"},
		Out:       []string{"error"},
	}

	receiver := "receiver"
	delegate := "delegate"

	expected := "return receiver.delegate.MethodName(string, int)"
	result := method.Delegate(receiver, delegate)

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
