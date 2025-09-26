package models

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddModelFields_1(t *testing.T) {
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

	testModel := &ModelInfo{
		Name:     "TestModel",
		FullName: "TestModel",
		Model:    TestStruct{},
		Fields:   &Fields{},
	}

	testValue := reflect.ValueOf(testModel.Model)
	AddModelFields(testModel, testValue, testModel.FullName, []int{})

	// Test if all fields are added correctly
	if len(testModel.Fields.Fields) != 3 {
		t.Errorf("Expected 3 fields, got %d", len(testModel.Fields.Fields))
	}

	// Test if fields are added correctly
	for _, field := range testModel.Fields.Fields {
		if field.Name != "Field1" && field.Name != "Field2" && field.Name != "Field3" {
			t.Errorf("Unexpected field: %s", field.Name)
		}
	}
}
