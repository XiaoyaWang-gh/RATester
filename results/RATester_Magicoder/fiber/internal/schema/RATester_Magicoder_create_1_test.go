package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func Testcreate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type testStruct struct {
		Field1 string
		Field2 int
		Field3 struct {
			Field4 string
		}
	}

	cache := &cache{
		m: make(map[reflect.Type]*structInfo),
	}

	info := cache.create(reflect.TypeOf(testStruct{}), "")

	// Test if the structInfo contains the correct number of fields
	if len(info.fields) != 3 {
		t.Errorf("Expected 3 fields, got %d", len(info.fields))
	}

	// Test if the structInfo contains the correct field names
	fieldNames := []string{"Field1", "Field2", "Field3"}
	for i, field := range info.fields {
		if field.name != fieldNames[i] {
			t.Errorf("Expected field name %s, got %s", fieldNames[i], field.name)
		}
	}

	// Test if the structInfo contains the correct field types
	fieldTypes := []reflect.Type{reflect.TypeOf(""), reflect.TypeOf(0), reflect.TypeOf(struct{}{})}
	for i, field := range info.fields {
		if field.typ != fieldTypes[i] {
			t.Errorf("Expected field type %s, got %s", fieldTypes[i], field.typ)
		}
	}
}
