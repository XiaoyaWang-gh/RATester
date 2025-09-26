package bean

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetConfig_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type TestStruct struct {
		Field1 string `beego:"field1"`
		Field2 int    `beego:"field2"`
	}

	testStruct := TestStruct{
		Field1: "test",
		Field2: 1,
	}

	beanValue := reflect.ValueOf(testStruct)

	tagAutoWireBeanFactory := &TagAutoWireBeanFactory{
		FieldTagParser: func(field reflect.StructField) *FieldMetadata {
			tag := field.Tag.Get("beego")
			return &FieldMetadata{
				DftValue: tag,
			}
		},
	}

	beanMetadata := tagAutoWireBeanFactory.getConfig(beanValue)

	if len(beanMetadata.Fields) != 2 {
		t.Errorf("Expected 2 fields, got %d", len(beanMetadata.Fields))
	}

	field1Metadata, ok := beanMetadata.Fields["Field1"]
	if !ok {
		t.Errorf("Expected Field1 in fields")
	}

	if field1Metadata.DftValue != "field1" {
		t.Errorf("Expected Field1 default value to be 'field1', got '%s'", field1Metadata.DftValue)
	}

	field2Metadata, ok := beanMetadata.Fields["Field2"]
	if !ok {
		t.Errorf("Expected Field2 in fields")
	}

	if field2Metadata.DftValue != "field2" {
		t.Errorf("Expected Field2 default value to be 'field2', got '%s'", field2Metadata.DftValue)
	}
}
