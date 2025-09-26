package hydratation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetStruct_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type TestStruct struct {
		Field1 string
		Field2 int
		Field3 bool
	}

	testStruct := TestStruct{
		Field1: "test",
		Field2: 123,
		Field3: true,
	}

	value := reflect.ValueOf(testStruct)
	err := setStruct(value)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}
