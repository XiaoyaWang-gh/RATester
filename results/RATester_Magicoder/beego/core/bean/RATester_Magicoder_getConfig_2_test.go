package bean

import (
	"fmt"
	"reflect"
	"testing"
)

func TestgetConfig_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type TestStruct struct {
		Field1 string `json:"field1"`
		Field2 int    `json:"field2"`
	}

	testStruct := TestStruct{
		Field1: "test",
		Field2: 123,
	}

	testStructValue := reflect.ValueOf(testStruct)

	tagAutoWireBeanFactory := &TagAutoWireBeanFactory{
		FieldTagParser: func(field reflect.StructField) *FieldMetadata {
			tag := field.Tag.Get("json")
			return &FieldMetadata{
				DftValue: tag,
			}
		},
	}

	beanMetadata := tagAutoWireBeanFactory.getConfig(testStructValue)

	if beanMetadata.Fields["Field1"].DftValue != "field1" {
		t.Errorf("Expected Field1 DftValue to be 'field1', but got %s", beanMetadata.Fields["Field1"].DftValue)
	}

	if beanMetadata.Fields["Field2"].DftValue != "field2" {
		t.Errorf("Expected Field2 DftValue to be 'field2', but got %s", beanMetadata.Fields["Field2"].DftValue)
	}
}
