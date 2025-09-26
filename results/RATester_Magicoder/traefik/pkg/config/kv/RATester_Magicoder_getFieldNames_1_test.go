package kv

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetFieldNames_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type TestStruct struct {
		Field1 string
		Field2 int
		Field3 bool
		Field4 struct {
			NestedField1 string
			NestedField2 int
		}
	}

	testStruct := TestStruct{
		Field1: "test",
		Field2: 123,
		Field3: true,
		Field4: struct {
			NestedField1 string
			NestedField2 int
		}{
			NestedField1: "nested",
			NestedField2: 456,
		},
	}

	expectedFieldNames := []string{
		"Field1",
		"Field2",
		"Field3",
		"Field4.NestedField1",
		"Field4.NestedField2",
	}

	fieldNames := getFieldNames("", reflect.TypeOf(testStruct))

	if !reflect.DeepEqual(fieldNames, expectedFieldNames) {
		t.Errorf("Expected field names: %v, but got: %v", expectedFieldNames, fieldNames)
	}
}
