package kv

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetRootFieldNames_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type TestStruct struct {
		Field1 string
		Field2 int
	}

	testStruct := TestStruct{
		Field1: "test",
		Field2: 123,
	}

	expected := []string{"Field1", "Field2"}
	result := getRootFieldNames("TestStruct", testStruct)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
